/*
Package secretmessages encrypts and decrypts a variety of algorithms.
*/

package main

import (
	"log"
	"net/http"

	"secretmessages/internal/pkg/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// Controllers
	index := controllers.NewIndex()

	// Routes
	router := httprouter.New()
	router.GET("/", index.Landing)
	router.POST("/", index.Convert)

	log.Fatal(http.ListenAndServe(":5000", router))
}
