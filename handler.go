package contact

import "github.com/ant0ine/go-json-rest/rest"

type Handler struct {
	provider Provider
}

func NewHandler(mp Provider) *Handler {
	return &Handler{provider: mp}
}

func (h *Handler) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	result, _ := h.provider.Get(id)
	w.WriteJson(result)
}

func (h *Handler) All(w rest.ResponseWriter, r *rest.Request) {
	result := h.provider.All()
	w.WriteJson(result)
}

func (h *Handler) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	h.provider.Delete(id)
	w.WriteJson(map[string]string{"Deleted": "Deleted"})
}

func (h *Handler) Update(w rest.ResponseWriter, r *rest.Request) {
	var i Information
	r.DecodeJsonPayload(&i)
	h.provider.Update(i)
	w.WriteJson(map[string]string{"Updateed": "Updated"})
}

func (h *Handler) Add(w rest.ResponseWriter, r *rest.Request) {
	var information Information
	r.DecodeJsonPayload(&information)
	err := h.provider.Add(&information)
	if err != nil {
		w.WriteJson(map[string]string{"Fail": "Fail to Add"})
	}
	w.WriteJson(map[string]string{"Added": "Added"})
}
