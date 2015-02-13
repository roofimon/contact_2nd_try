package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
    "contact"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := contact.MakeRestRouter()

	if err != nil {
		log.Fatal(err)
	}
    
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
