package router

import (
	"go_server_practice/server/controller"
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
		controller.Index,
	},
	Route{
		"Create",
		"POST",
		"/create",
		controller.CreatePerson,
	},
	Route{
		"Read",
		"GET",
		"/read/{userId}",
		controller.ReadPerson,
	},
	Route{
		"ReadAll",
		"GET",
		"/read",
		controller.ReadAllPerson,
	},
	Route{
		"Count",
		"GET",
		"/count",
		controller.CountAllPerson,
	},
	Route{
		"Update",
		"POST",
		"/update",
		controller.UpdatePerson,
	},
	Route{
		"Delete",
		"GET",
		"/delete/{userId}",
		controller.DeletePerson,
	},
	Route{
		"DeleteAll",
		"GET",
		"/clear",
		controller.DeleteAllPerson,
	},
	Route{
		"Error",
		"GET",
		"/error",
		controller.ErrorTest,
	},
}
