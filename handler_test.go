package contact

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"io/ioutil"
	"log"
	"testing"
)

func initialHandler() rest.ResourceHandler {
	contactHandler := rest.ResourceHandler{
		DisableJsonIndent: true,
		ErrorLogger:       log.New(ioutil.Discard, "", 0),
	}
	contact := NewRouter()
	contactHandler.SetRoutes(contact.All, contact.Get)
	return contactHandler
}

func TestAll(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("GET", "http://1.2.3.4/contact", nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"All":"All"}`)
}

func TestGet(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, &contactHandler, test.MakeSimpleRequest("GET", "http://1.2.3.4/contact/123", nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{"Id":"123"}`)
}
