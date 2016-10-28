package main

import (
	"fmt"
	"strings"
	"net/http"

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
	defer close()

	// mc := queryAll()
	fmt.Fprintln(w, "JSON output of full db")
}

func AllSearch(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	searchStr := varsIn["searchStr"]
	fmt.Fprintln(w, "JSON output of matching db entries starting with: ", searchStr)
}

func AllType(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	typeStr := varsIn["typeStr"]
	fmt.Fprintln(w, "JSON output of matching db entries with type: ", typeStr)
}

func Message(w http.ResponseWriter, r *http.Request) {
	varsIn := mux.Vars(r)
	id := varsIn["id"]
	vars := varsIn["vars"]
	fmt.Fprintln(w, "MessageId: ", id)
	fmt.Fprintln(w, "Vars Length: ", len(vars))
	if len(vars) > 0 {
		splitVars := strings.Split(strings.Replace(vars, " ", "", -1), ",")
		fmt.Fprintln(w, "Return message with variables substituted ", splitVars)
	}
}
