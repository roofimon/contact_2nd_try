package contact

import "github.com/ant0ine/go-json-rest/rest"

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	w.WriteJson(map[string]string{"Id": id})
}

func (h *Handler) All(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"All": "All"})
}
