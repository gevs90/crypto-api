package routes

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
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", func(w http.ResponseWriter, r *http.Request) {}},
	Route{"Index", "POST", "/v1/encrypt", Encrypt},
	Route{"Index", "GET", "/v1/encrypt/{record_id}", Decrypt},
}
