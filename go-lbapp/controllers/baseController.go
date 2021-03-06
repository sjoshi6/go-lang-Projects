package controllers

import (
	"go-lbapp/api"
	"net/http"

	"github.com/gorilla/mux"
)

// Central router for all API requests

var routes = Routes{
	Route{
		"signup",
		"POST",
		"/v1/signup",
		api.CreateAccount,
	},
	Route{
		"login",
		"POST",
		"/v1/login",
		api.ConfirmCredentials,
	},
	Route{
		"search_events",
		"GET",
		"/v1/search_events",
		api.SearchEventsByRange,
	},
	Route{
		"subscribe",
		"POST",
		"/v1/subscribe/{eventid}",
		api.JoinEvent,
	},
	Route{
		"unsubscribe",
		"POST",
		"/v1/unsubscribe/{eventid}",
		api.LeaveEvent,
	},
}

// SetBaseRoutes : Get an object of gorilla mux router
func SetBaseRoutes(router *mux.Router) *mux.Router {

	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = APILogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
