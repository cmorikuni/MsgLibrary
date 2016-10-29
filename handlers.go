package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Fprintln(w, "/all")
	fmt.Fprintln(w, "/id/{id int}")
	fmt.Fprintln(w, "/str/{search string}")
	fmt.Fprintln(w, "/type/{type string}")
	fmt.Fprintln(w, "/message/{id int}")
	fmt.Fprintln(w, "/message/{id int}/[var1,var2,var3]")
	fmt.Fprintln(w, "/admin/insert/{type string}/{msg string} = need to use %25 for percent")
	fmt.Fprintln(w, "/admin/delete/{id int}")
	fmt.Fprintln(w, "/admin/update/{id}/{field}/{value}")
}

func All(w http.ResponseWriter, r *http.Request) {
	open()
	mc := queryAll()
	fmt.Fprintln(w, toJson(mc))
}

func AllId(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	searchId := varsIn["id"]

	open()
	mc := queryId(searchId)
	fmt.Fprintln(w, toJson(mc))
}

func AllSearch(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	searchStr := varsIn["searchStr"]

	open()
	mc := querySearch(searchStr)
	fmt.Fprintln(w, toJson(mc))
}

func AllType(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	typeStr := varsIn["typeStr"]
	
	open()
	mc := queryType(typeStr)
	fmt.Fprintln(w, toJson(mc))
}

func Message(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	id := varsIn["id"]
	vars := varsIn["vars"]
	
	open()
	var splitVars []string
	if len(vars) > 0 {
		splitVars = strings.Split(strings.Replace(vars, " ", "", -1), ",")
	}
	fmt.Fprintln(w, queryMsg(id, splitVars))
}

func Insert(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	typeStr := varsIn["typeStr"]
	msg := varsIn["msg"]

	open()
	mc := insert(typeStr, msg)
	fmt.Fprintln(w, toJson(mc))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	id := varsIn["id"]

	open()
	fmt.Fprintln(w, delete(id))
}

func Update(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	id := varsIn["id"]
	field := varsIn["field"]
	value := varsIn["value"]

	open()
	mc := update(id, field, value)
	fmt.Fprintln(w, toJson(mc))
}

func toJson(mc []*MsgCategory) string {
   bytes, err := json.MarshalIndent(mc, "", "  ")
   if err != nil {
                   fmt.Println(err.Error())
                   os.Exit(1)
   }
   return string(bytes)
}