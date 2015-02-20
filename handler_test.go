package contact

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

var content *map[string]string = &map[string]string{"Id": "d9356b78-6b54-4391-80d0-af2c4949d973", "Email": "first@email.com", "Title": "First", "Content": "First Content"}
var infomation = map[string]string{"Id": "d9356b78-6b54-4391-80d0-af2c4949d973", "Email": "first@email.com", "Title": "First", "Content": "First Content"}
var ID = "d9356b78-6b54-4391-80d0-af2c4949d973"
var HOST = "http://1.2.3.4/contact"
var CONTACT_WITH_ID = HOST + "/" + ID

func initialHandler() http.Handler {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, _ := MakeRestRouter(NewStubProvider())
	api.SetApp(router)
	return api.MakeHandler()
}

func TestAdd(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, contactHandler, test.MakeSimpleRequest("POST", HOST, content))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}

func TestAll(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, contactHandler, test.MakeSimpleRequest("GET", HOST, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	var body []map[string]string
	expect := []map[string]string{infomation}
	recorded.DecodeJsonPayload(&body)
	if !reflect.DeepEqual(body, expect) {
		t.Errorf("Body %v expected got %v", expect, body)
	}
}

func TestGet(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, contactHandler, test.MakeSimpleRequest("GET", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	var body map[string]string
	recorded.DecodeJsonPayload(&body)
	if !reflect.DeepEqual(body, infomation) {
		t.Errorf("Body %v expected got %v", infomation, body)
	}
}

func TestUpdate(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, contactHandler, test.MakeSimpleRequest("PUT", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}

func TestDelete(t *testing.T) {
	contactHandler := initialHandler()

	recorded := test.RunRequest(t, contactHandler, test.MakeSimpleRequest("DELETE", CONTACT_WITH_ID, nil))

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
