package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID     string
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Persons []Person

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/person", GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Jason, golang rock!")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	persons := Persons{
		Person{
			ID:     "123",
			Name:   "Jason",
			Age:    27,
			Gender: "m",
		},
	}

	json.NewEncoder(w).Encode(persons)
}
