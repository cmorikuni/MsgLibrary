package main

import (
	"strings"
	"strconv"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Msg struct {
	Oid     int    `json:"oid"`
	Message string `json:"message"`
	Called  int    `json:"called"`
}

type MsgCategory struct {
	Category string `json:"category"`
	Messages []Msg  `json:"messages"`
}

func open() {
	var err error
	db, err = sql.Open("sqlite3", "./msgLibrary.db")
	checkErr(err)
}

func insert(typeStr string, msg string) []*MsgCategory {
	stmt, err := db.Prepare("INSERT INTO messages(category, message, called) VALUES (?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(strings.ToUpper(typeStr), msg, 0)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return queryId(strconv.Itoa(int(id)))
}

func delete(id string) string {
	stmt, err := db.Prepare("DELETE FROM messages WHERE oid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return strconv.Itoa(int(affect)) + " rows deleted."
}

func update(id string, field string, value string) []*MsgCategory {
	stmt, err := db.Prepare("UPDATE messages SET " + field + "=? WHERE oid=?")
	checkErr(err)

	_, err := stmt.Exec(value, id)
	checkErr(err)

	return queryId(id)
}

func queryAll() []*MsgCategory {
	rows, err := db.Query("SELECT * FROM messages")
	checkErr(err)
	defer rows.Close()

	return queryPrint(rows)
}

func queryId(id string) []*MsgCategory {
	rows, err := db.Query("SELECT * FROM messages WHERE oid == " + id)
	checkErr(err)
	defer rows.Close()

	return queryPrint(rows)
}

func querySearch(str string) []*MsgCategory {
	rows, err := db.Query("SELECT * FROM messages WHERE message LIKE '%" + str + "%'")
	checkErr(err)
	defer rows.Close()

	return queryPrint(rows)
}

func queryType(str string) []*MsgCategory {
	rows, err := db.Query("SELECT * FROM messages WHERE category == '" + strings.ToUpper(str) + "'")
	checkErr(err)
	defer rows.Close()

	return queryPrint(rows)
}

func queryMsg(id string, vars []string) string {
	rows, err := db.Query("SELECT oid, message, called FROM messages WHERE oid == " + id)
	checkErr(err)
	defer rows.Close()

	// If vars replace %% token
	return msgPrint(rows, vars)
}

func queryPrint(rows *sql.Rows) []*MsgCategory {
	mcs := []*MsgCategory{}
	for rows.Next() {
		var oid int
		var category string
		var message string
		var called int

		err := rows.Scan(&oid, &category, &message, &called)
		checkErr(err)

		m := Msg{oid, message, called}
		ind := categoryExists(category, mcs)
		if ind == -1 {
			mc := new(MsgCategory)
			mc.Category = category

			mcs = append(mcs, mc)
			ind = len(mcs)-1
		}
		mcs[ind].Messages = append(mcs[ind].Messages, m)
	}
	return mcs
}

func msgPrint(rows *sql.Rows, vars []string) string {
	var oid int
	var message string
	var called int
	for rows.Next() {
		err := rows.Scan(&oid, &message, &called)
		checkErr(err)

		cnt := strings.Count(message, "%%")
		if len(vars) != cnt {
			return "ERROR: variables (" + strconv.Itoa(len(vars)) + ") and token counts (" + strconv.Itoa(cnt) + ") do not match."
		}
		
		for _, v := range vars {
			message = strings.Replace(message, "%%", v, 1)
		}
	}

	called++
	updateCalled(oid, called)
	return message
}

func categoryExists(cat string, mc []*MsgCategory) int {
	for ind, item := range mc {
		if item.Category == cat {
			return ind
		}
	}
	return -1
}

func updateCalled(oid int, called int) {
    stmt, err := db.Prepare("UPDATE messages SET called=? WHERE oid=?")
    checkErr(err)

    _, err = stmt.Exec(called, oid)
    checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
