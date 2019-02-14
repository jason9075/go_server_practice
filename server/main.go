package main

import (
	"go_server_practice/server/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
