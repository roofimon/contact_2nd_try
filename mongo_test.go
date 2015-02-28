package contact

import (
	"os"
	"reflect"
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

func TestGetAllContact(t *testing.T) {
	c := s.DB("test").C("contact")
	defer c.DropCollection()

	var expectInfos = []Information{
		{Email: "mail@email.com", Title: "Information", Content: "your content"},
		{Email: "mail1@email.com", Title: "Information1", Content: "your content"},
		{Email: "mail2@email.com", Title: "Information2", Content: "your content"},
	}
	c.Insert(expectInfos[0])
	c.Insert(expectInfos[1])
	c.Insert(expectInfos[2])

	mp := &MongoProvider{session: s}

	var infos []Information = mp.All()

	if !reflect.DeepEqual(infos, expectInfos) {
		t.Error("expect infomations is", expectInfos, "but is", infos)
	}
}
