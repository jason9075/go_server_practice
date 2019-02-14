package controller

import (
	"encoding/json"
	"fmt"
	"go_server_practice/server/model"
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

	var person model.Person
	json.Unmarshal(body, &person)

	newID := model.InsertPerson(person)

	w.Write(msgOk(newID))

}

func ReadPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	person := model.QueryPerson(userID)

	json.NewEncoder(w).Encode(person)

}

func ReadAllPerson(w http.ResponseWriter, r *http.Request) {
	persons := model.ReadAllPerson()

	json.NewEncoder(w).Encode(persons)

}

func CountAllPerson(w http.ResponseWriter, r *http.Request) {
	model.PrintAllPerson()
	persons := model.ReadAllPerson()

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

	isSuccess := model.UpdatePerson(userID, payload)

	if isSuccess {
		w.Write(msgOk("done"))
		return
	}
	w.Write(msgErr("update fail..., ID not exist!"))

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	isSuccess := model.DeletePerson(userID)

	if isSuccess {
		w.Write(msgOk("done"))
		return
	}
	w.Write(msgErr("ID not exist!"))
}

func DeleteAllPerson(w http.ResponseWriter, r *http.Request) {
	model.DeleteAllPerson()
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
