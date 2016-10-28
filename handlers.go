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
	fmt.Fprintln(w, "/all/{search string}")
	fmt.Fprintln(w, "/all/{type string}")
	fmt.Fprintln(w, "/message/id")
	fmt.Fprintln(w, "/message/id/[var1,var2,var3]")
}

func All(w http.ResponseWriter, r *http.Request) {
	open()
	mc := queryAll()
	fmt.Fprintln(w, toJson(mc))
}

func AllSearch(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	searchStr := varsIn["searchStr"]
	mc := querySearch(searchStr)
	fmt.Fprintln(w, toJson(mc))
}

func AllType(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	typeStr := varsIn["typeStr"]
	mc := queryType(typeStr)
	fmt.Fprintln(w, toJson(mc))
}

func Message(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	id := varsIn["id"]
	vars := varsIn["vars"]
	
	var splitVars []string
	if len(vars) > 0 {
		splitVars = strings.Split(strings.Replace(vars, " ", "", -1), ",")
	}
	fmt.Fprintln(w, queryMsg(id, splitVars))
}

func toJson(mc []*MsgCategory) string {
   bytes, err := json.Marshal(mc)
   if err != nil {
                   fmt.Println(err.Error())
                   os.Exit(1)
   }
   return string(bytes)
}