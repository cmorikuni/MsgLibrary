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
		"AllId",
		"GET",
		"/id/{id}",
		AllId,
	},
	Route{
		"AllSearch",
		"GET",
		"/str/{searchStr}",
		AllSearch,
	},
	Route{
		"AllType",
		"GET",
		"/type/{typeStr}",
		AllType,
	},
	Route{
		"Message",
		"GET",
		"/message/{id}",
		Message,
	},
	Route{
		"Message",
		"GET",
		"/message/{id}/{vars}",
		Message,
	},
	Route{
		"AdminInsert",
		"GET",
		"/admin/insert/{typeStr}/{msg}",
		Insert,
	},
	Route{
		"AdminDelete",
		"GET",
		"/admin/delete/{id}",
		Delete,
	},
	Route{
		"AdminUpdate",
		"GET",
		"/admin/update/{id}/{field}/{value}",
		Update,
	},
}
