package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// func NewRouter() *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		var handler http.Handler
// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)
// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(route.HandlerFunc)
// 	}
//
// 	return router
// }

// var routes = Routes{
// 	Route{
// 		Name:        "Index",
// 		Method:      "GET",
// 		Pattern:     "/",
// 		HandlerFunc: Index,
// 	},
// 	Route{
// 		Name:        "HeroIndex",
// 		Method:      "GET",
// 		Pattern:     "/heroes",
// 		HandlerFunc: HeroIndex,
// 	},
// 	Route{
// 		Name:        "HeroShow",
// 		Method:      "GET",
// 		Pattern:     "/heroes/{HeroId}",
// 		HandlerFunc: HeroShow,
// 	},
// }
