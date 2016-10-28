package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"All",
		"GET",
		"/all",
		All,
	},
	Route{
		"AllSearch",
		"GET",
		"/all/{searchStr}",
		AllSearch,
	},
	Route{
		"AllType",
		"GET",
		"/all/{typeStr}",
		AllType,
	},
	Route{
		"Message",
		"POST",
		"/message/{id}",
		Message,
	},
	Route{
		"Message",
		"POST",
		"/message/{id}/{vars}",
		Message,
	},
}
