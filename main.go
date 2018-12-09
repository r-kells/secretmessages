/*
Package secretmessages encrypts and decrypts a variety of algorithms.
*/

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/r-kells/secretmessages/controllers"
)

func main() {

	// Controllers
	index := controllers.NewIndex()

	// Routes
	router := httprouter.New()
	router.GET("/", index.Landing)
	router.POST("/", index.Convert)

	log.Fatal(http.ListenAndServe(":8080", router))
}
