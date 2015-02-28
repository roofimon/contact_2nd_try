package contact

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type MongoProvider struct {
	session *mgo.Session
}

func NewMongoProvider() *MongoProvider {
	session, _ := mgo.Dial("localhost")
	return &MongoProvider{session}
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
	v := bson.M{"_id": bson.ObjectIdHex(id)}
	err = do(mp.session.Clone(), find, v).one(&result)
	err = handleError(err)

	return
}

func (mp *MongoProvider) All() (result []Information) {
	do(mp.session.Clone(), find, nil).all(&result)
	return
}

type query struct {
	s *mgo.Session
	q chan *mgo.Query
}

func (q *query) all(v interface{}) error {
	defer q.s.Close()
	return q.retrieveQuery().All(v)
}

func (q *query) one(v interface{}) error {
	defer q.s.Close()
	return q.retrieveQuery().One(v)
}

func (q *query) retrieveQuery() *mgo.Query {
	for {
		select {
		case query := <-q.q:
			return query
		}
	}
}

type queryFunc func(c *mgo.Collection, v interface{}) *mgo.Query

func do(s *mgo.Session, fn queryFunc, v interface{}) *query {
	q := make(chan *mgo.Query)
	go func() {
		q <- fn(ContactCollection(s), v)
	}()
	return &query{s, q}
}

func find(c *mgo.Collection, v interface{}) *mgo.Query {
	return c.Find(v)
}

func (mp *MongoProvider) Update(i Information) error {
	err := sessionHandler(update)(i)
	return handleError(err)
}

func (mp *MongoProvider) Delete(id string) error {
	err := sessionHandler(delete)(id)
	return handleError(err)
}

func (mp *MongoProvider) Add(i *Information) error {
	err := sessionHandler(add)(i)
	return handleError(err)
}

type action func(*mgo.Collection, interface{}) error
type wrapper func(interface{}) error

func sessionHandler(a action) wrapper {
	return func(input interface{}) error {
		s := CloneSession()
		c := ContactCollection(s)
		defer s.Close()

		return a(c, input)
	}
}

func update(c *mgo.Collection, input interface{}) error {
	i, _ := input.(Information)
	target := bson.M{"id": i.Id}
	change := bson.M{"$set": bson.M{"id": i.Id, "email": i.Email, "title": i.Title, "content": i.Content}}
	return c.Update(target, change)
}

func delete(c *mgo.Collection, input interface{}) error {
	id, _ := input.(string)
	target := bson.M{"id": id}
	return c.Remove(target)
}

func add(c *mgo.Collection, input interface{}) error {
	i, _ := input.(Information)
	return c.Insert(i)
}

func handleError(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
