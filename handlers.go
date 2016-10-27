package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Fprintln(w, "/all")
	fmt.Fprintln(w, "/all/{search string}")
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
	vars := mux.Vars(r)
	searchStr := vars["searchStr"]
	fmt.Fprintln(w, "JSON output of matching db entries starting with: ", searchStr)
}

func Message(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id := vars["id"]
	fmt.Fprintln(w, "Return message with variables substituted", vars)
}
