package contact

import (
	"os"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var s *mgo.Session

func TestMain(m *testing.M) {
	s, _ = mgo.Dial("localhost")
	defer s.Close()

	code := m.Run()

	os.Exit(code)
}

func TestGetID(t *testing.T) {
	c := s.DB("test").C("contact")
	defer c.DropCollection()

	id := bson.NewObjectId()
	c.Insert(bson.M{
		"_id":     id,
		"email":   "mail@email.com",
		"title":   "Information",
		"content": "your content",
	})

	mp := &MongoProvider{session: s}

	info, _ := mp.Get(id.Hex())

	if info.Email != "mail@email.com" {
		t.Error("expect email : mail@email.com")
	}
	if info.Title != "Information" {
		t.Error("expect title: Information")
	}
	if info.Content != "your content" {
		t.Error("expect content: your content")
	}
}
