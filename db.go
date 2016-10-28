package main

// import (
//            "encoding/json"
//            "fmt"
//            "io/ioutil"
//            "os"
//            "bytes"
// )

// type MsgCategory []struct {
//            Category string   `json:"category"`
//            Messages []string `json:"messages"`
// }

// func (mc MsgCategory) toString() string {
//            bytes, err := json.Marshal(mc)
//            if err != nil {
//                            fmt.Println(err.Error())
//                            os.Exit(1)
//            }
//            return string(bytes)
// }

// func getMessages() MsgCategory {
//            raw, err := ioutil.ReadFile("./msgs.json")
//            if err != nil {
//                            fmt.Println(err.Error())
//                            os.Exit(1)
//            }
//            dec := json.NewDecoder(bytes.NewReader(raw))
//            var mc MsgCategory
//            dec.Decode(&mc)
//            return mc
// }

// [
//   {
//      category: string,
//      messages: [
//          string
//      ]
//   }
//]

// CREATE TABLE messages(oid INTEGER PRIMARY KEY, category TEXT, message TEXT, called INTEGER)
// INSERT INTO messages(category, message, called) VALUES ('random', 'Invalid %% message format needs to be %%.', 0)
// QUERY SELECT: SELECT * FROM messages WHERE message LIKE '%test msg%'

import (
	"fmt"
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

func categoryExists(cat string, mc []*MsgCategory) int {
	for ind, item := range mc {
		if item.Category == cat {
			return ind
		}
	}
	return -1
}

func queryAll() []*MsgCategory {
	rows, err := db.Query("SELECT * FROM messages")
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
	rows, err := db.Query("SELECT * FROM messages WHERE category == '" + str + "'")
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
		fmt.Println(cnt)
		
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

