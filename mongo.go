package contact

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type MongoProvider struct{}

func NewMongoProvider() *MongoProvider {
	session, _ = mgo.Dial("localhost")
	return &MongoProvider{}
}

func CloneSession() *mgo.Session {
	s := session.Clone()
	s.SetMode(mgo.Monotonic, true)
	return s
}

func ContactCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("test").C("contact")
}

func (mp *MongoProvider) Get(id string) (result Information, err error) {
	s := CloneSession()
	defer s.Close()
	c := ContactCollection(s)

	err = c.Find(bson.M{"id": id}).One(&result)
	err = handleError(err)

	return
}

func (mp *MongoProvider) All() []Information {
	result := []Information{}
	s := CloneSession()
	c := ContactCollection(s)

	c.Find(nil).All(&result)
	return result
}

func (mp *MongoProvider) Update(i Information) error {
	target := bson.M{"id": i.Id}
	change := bson.M{"$set": bson.M{"id": i.Id, "email": i.Email, "title": i.Title, "content": i.Content}}

	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	err := c.Update(target, change)
	return handleError(err)
}

func (mp *MongoProvider) Delete(id string) error {
	target := bson.M{"id": id}

	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()

	err := c.Remove(target)
	return handleError(err)
}

func (mp *MongoProvider) Add(i *Information) error {
	s := CloneSession()
	c := ContactCollection(s)
	defer s.Close()
	err := c.Insert(i)
	return handleError(err)
}

func handleError(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
