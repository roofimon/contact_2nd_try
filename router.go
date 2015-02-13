package contact

import "github.com/ant0ine/go-json-rest/rest"

type Router struct {
	All, Get *rest.Route
}

func NewRouter() *Router {
	handler := NewHandler()
	All := &rest.Route{"GET", "/contact", handler.All}
	Get := &rest.Route{"GET", "/contact/:id", handler.Get}
	return &Router{All: All, Get: Get}
}

func MakeRestRouter() (rest.App, error) {
	contact := NewRouter()

	return rest.MakeRouter(
		contact.All,
		contact.Get,
	)
}
