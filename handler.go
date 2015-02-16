package contact

import (
	"github.com/ant0ine/go-json-rest/rest"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

type Information struct {
	Id      string
	Email   string
	Title   string
	Content string
}

var session *mgo.Session

func (h *Handler) Get(w rest.ResponseWriter, r *rest.Request) {
	session, _ = mgo.Dial("localhost")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	contact := session.DB("test").C("contact")
	id := r.PathParam("id")
	result := Information{}
	contact.Find(bson.M{"id": id}).One(&result)
	w.WriteJson(result)
}

func (h *Handler) All(w rest.ResponseWriter, r *rest.Request) {
	session, _ = mgo.Dial("localhost")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	contact := session.DB("test").C("contact")
	result := []Information{}
	contact.Find(nil).All(&result)
	w.WriteJson(result)
}

func (h *Handler) Delete(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Deleted": "Deleted"})
}

func (h *Handler) Update(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Updateed": "Updated"})
}

func (h *Handler) Add(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Added": "Added"})
}
