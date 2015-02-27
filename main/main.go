package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/roofimon/contact"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	mp := contact.NewMongoProvider()
	router, err := contact.MakeRestRouter(mp)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
