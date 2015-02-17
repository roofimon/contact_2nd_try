package contact

import (
	"github.com/ant0ine/go-json-rest/rest"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Information struct {
	Id      string
	Email   string
	Title   string
	Content string
}

type Handler struct{}

var session *mgo.Session

func NewHandler() *Handler {
	session, _ = mgo.Dial("localhost")
	return &Handler{}
}

func CloneSession() *mgo.Session {
	s := session.Clone()
	s.SetMode(mgo.Monotonic, true)
	return s
}

func ContactCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("test").C("contact")
}

func (h *Handler) Get(w rest.ResponseWriter, r *rest.Request) {
	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	id := r.PathParam("id")
	result := Information{}
	c.Find(bson.M{"id": id}).One(&result)
	w.WriteJson(result)
}

func (h *Handler) All(w rest.ResponseWriter, r *rest.Request) {
	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	result := []Information{}
	c.Find(nil).All(&result)
	w.WriteJson(result)
}

func (h *Handler) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	target := bson.M{"id": id}

	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	err := c.Remove(target)
	if err != nil {
	}

	w.WriteJson(map[string]string{"Deleted": "Deleted"})
}

func (h *Handler) Update(w rest.ResponseWriter, r *rest.Request) {
	var i Information
	r.DecodeJsonPayload(&i)
	target := bson.M{"id": i.Id}
	change := bson.M{"$set": bson.M{"id": i.Id, "email": i.Email, "title": i.Title, "content": i.Content}}

	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	err := c.Update(target, change)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(map[string]string{"Updateed": "Updated"})
}

func (h *Handler) Add(w rest.ResponseWriter, r *rest.Request) {
	var information Information
	r.DecodeJsonPayload(&information)

	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()
	err := c.Insert(information)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteJson(map[string]string{"Added": "Added"})
}
