package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodosIndex", "GET", "/todos", TodosIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
	Route{"TodoDelete", "DELETE", "/todos/{todoId}", TodoDelete},
}
