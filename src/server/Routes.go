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
		"Home",
		"GET",
		"/",
		Index,
	},
	Route{
		"Create",
		"POST",
		"/create",
		CreatePerson,
	},
	Route{
		"Read",
		"GET",
		"/read/{userId}",
		ReadPerson,
	},
	Route{
		"ReadAll",
		"GET",
		"/read",
		ReadAllPerson,
	},
	Route{
		"Count",
		"GET",
		"/count",
		CountAllPerson,
	},
	Route{
		"Update",
		"POST",
		"/update",
		UpdatePerson,
	},
	Route{
		"Delete",
		"GET",
		"/delete/{userId}",
		DeletePerson,
	},
	Route{
		"Error",
		"GET",
		"/error",
		ErrorTest,
	},
}
