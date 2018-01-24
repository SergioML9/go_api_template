package routes

import (
	"net/http"
	"../delivery"
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
		"/content",
		delivery.GetContent,
	},
}