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

type testContactFunc func(t *testing.T, c *mgo.Collection, mp *MongoProvider)

func TestMongoProvider(t *testing.T) {
	mp := &MongoProvider{session: s}

	for _, fn := range []testContactFunc{testGetContactByID, testGetAllContact} {
		c := s.DB("test").C("contact")
		fn(t, c, mp)
		c.DropCollection()
	}
}

func testGetContactByID(t *testing.T, c *mgo.Collection, mp *MongoProvider) {
	id := bson.NewObjectId()
	c.Insert(bson.M{
		"_id":     id,
		"email":   "mail@email.com",
		"title":   "Information",
		"content": "your content",
	})

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

func testGetAllContact(t *testing.T, c *mgo.Collection, mp *MongoProvider) {
	var expectInfos = []Information{
		{Email: "mail@email.com", Title: "Information", Content: "your content"},
		{Email: "mail1@email.com", Title: "Information1", Content: "your content"},
		{Email: "mail2@email.com", Title: "Information2", Content: "your content"},
	}
	c.Insert(expectInfos[0])
	c.Insert(expectInfos[1])
	c.Insert(expectInfos[2])

	var infos []Information = mp.All()

	if !reflect.DeepEqual(infos, expectInfos) {
		t.Error("expect infomations is", expectInfos, "but is", infos)
	}
}
