package main

import (
	"hello/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
