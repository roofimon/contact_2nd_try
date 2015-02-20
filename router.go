package contact

import "github.com/ant0ine/go-json-rest/rest"

type Router struct {
	All, Get, Delete, Update, Add *rest.Route
}

func NewRouter(mp Provider) *Router {
	handler := NewHandler(mp)

	All := &rest.Route{
		HttpMethod: "GET",
		PathExp:    "/contact",
		Func:       handler.All,
	}

	Get := &rest.Route{
		HttpMethod: "GET",
		PathExp:    "/contact/:id",
		Func:       handler.Get,
	}

	Delete := &rest.Route{
		HttpMethod: "DELETE",
		PathExp:    "/contact/:id",
		Func:       handler.Delete,
	}

	Update := &rest.Route{
		HttpMethod: "PUT",
		PathExp:    "/contact/:id",
		Func:       handler.Get,
	}

	Add := &rest.Route{
		HttpMethod: "POST",
		PathExp:    "/contact",
		Func:       handler.Add,
	}

	return &Router{All: All, Get: Get, Delete: Delete, Update: Update, Add: Add}
}

func MakeRestRouter(mp Provider) (rest.App, error) {
	contact := NewRouter(mp)

	return rest.MakeRouter(
		contact.All,
		contact.Get,
		contact.Delete,
		contact.Update,
		contact.Add,
	)
}
