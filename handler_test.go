package contact

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"io/ioutil"
	"log"
	"testing"
)

var content *map[string]string = &map[string]string{"Id": "d9356b78-6b54-4391-80d0-af2c4949d973", "Email": "first@email.com", "Title": "First", "Content": "First Content"}
var information = `{"Id":"d9356b78-6b54-4391-80d0-af2c4949d973","Email":"first@email.com","Title":"First","Content":"First Content"}`
var ID = "d9356b78-6b54-4391-80d0-af2c4949d973"
var HOST = "http://1.2.3.4/contact"
var CONTACT_WITH_ID = HOST + "/" + ID

func initialHandler() rest.ResourceHandler {
	contactHandler := rest.ResourceHandler{
		DisableJsonIndent: true,
		ErrorLogger:       log.New(ioutil.Discard, "", 0),
	}
	contact := NewRouter(NewStubProvider())
	contactHandler.SetRoutes(contact.All, contact.Get, contact.Delete, contact.Update, contact.Add)
	return contactHandler
}

func TestAdd(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("POST", HOST, content))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}

func TestAll(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("GET", HOST, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`[` + information + `]`)
}

func TestGet(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("GET", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(information)
}

func TestUpdate(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("PUT", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}

func TestDelete(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("DELETE", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
