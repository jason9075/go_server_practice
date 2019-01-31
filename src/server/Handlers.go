package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Jason, golang rock!")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(msgErr("Error reading request body"))
		return
	}

	var person Person
	json.Unmarshal(body, &person)

	db := GetDb()
	newID := db.Insert(person)
	db.PrintAll()

	w.Write(msgOk(newID))

}

func ReadPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	db := GetDb()
	person := db.Query(userID)

	json.NewEncoder(w).Encode(person)

}

func ReadAllPerson(w http.ResponseWriter, r *http.Request) {
	db := GetDb()
	persons := db.ReadAll()

	json.NewEncoder(w).Encode(persons)

}

func CountAllPerson(w http.ResponseWriter, r *http.Request) {
	db := GetDb()
	persons := db.ReadAll()

	var res = map[string]interface{}{"status": "success!", "msg": len(persons)}
	response, _ := json.Marshal(res)
	w.Write(response)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.Write(msgErr("Error decoding payload"))
		return
	}

	userID := payload["user_id"].(string)
	if userID == "" {
		w.Write(msgErr("Error user_id can't not be nil!"))
		return
	}

	db := GetDb()
	isSuccess := db.Update(userID, payload)

	if isSuccess {
		w.Write(msgOk("done"))
		return
	}
	w.Write(msgErr("update fail..., ID not exist!"))

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	db := GetDb()
	isSuccess := db.Delete(userID)

	if isSuccess {
		w.Write(msgOk("done"))
		return
	}
	w.Write(msgErr("ID not exist!"))
}

func ErrorTest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error reading request body",
		http.StatusInternalServerError)
}

func msgOk(msg string) []byte {
	var res = map[string]string{"status": "success!", "msg": msg}
	response, _ := json.Marshal(res)
	return response
}

func msgErr(msg string) []byte {
	var res = map[string]string{"status": "fail", "msg": msg}
	response, _ := json.Marshal(res)
	return response
}
