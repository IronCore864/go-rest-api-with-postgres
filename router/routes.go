package router

import (
	"hello/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"New",
		"POST",
		"/new",
		handler.New,
	},
	Route{
		"Quote",
		"GET",
		"/quote",
		handler.Quote,
	},
}
