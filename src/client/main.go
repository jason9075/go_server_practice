package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const DomainName string = "http://localhost:8080/"

const APICreate string = DomainName + "create"
const APIRead string = DomainName + "read"
const APICount string = DomainName + "count"
const APIUpdate string = DomainName + "update"
const APIDelete string = DomainName + "delete"
const APIDeleteAll string = DomainName + "clear"

func main() {
	// Basic CRUD
	fmt.Println("## Basic CRUD ##")
	personID, _ := requestCreate("Jason", 25, "m", true)
	requestRead(personID)
	requestUpdate(personID, "Jason_edit", 27, "m")
	requestRead(personID)
	requestDump()
	requestDelete(personID)

	total := 5000

	// For loop request
	start := time.Now()

	for i := 0; i < total; i++ {
		_, _ = requestCreate("Jason", i, "m", false)
	}

	elapsed := time.Since(start)
	fmt.Printf("for-loop %d requests took %s\n", total, elapsed)

	requestDeleteAll()

	// Goroutine
	fmt.Println("## Goroutine ##")

	batchSize := 8
	requestResults := make(chan string)
	requestBatch := make(chan int, batchSize)

	start = time.Now()

	for i := 0; i < total; i++ {
		go func(index int) {
			requestBatch <- index
			personID, _ := requestCreate("Jason", index, "m", false)
			<-requestBatch
			requestResults <- personID
		}(i)
	}

	for i := 0; i < total; i++ {
		<-requestResults
	}

	elapsed = time.Since(start)
	fmt.Printf("goroutine %d requests took %s\n", total, elapsed)

	requestCount()

}

func requestCreate(name string, age int, gender string, verbose bool) (string, error) {
	jsonData := map[string]interface{}{"name": name, "age": age, "gender": gender}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(APICreate, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "", errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	result := make(map[string]string)
	err = json.Unmarshal(data, &result)
	if verbose {
		println("Create: " + result["msg"])
	}
	return result["msg"], nil

}

func requestRead(pID string) error {
	response, err := http.Get(APIRead + "/" + pID)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	println("Read: " + string(data))
	return nil
}

func requestDump() error {
	response, err := http.Get(APIRead)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	println("Dump: " + string(data))
	return nil
}

func requestCount() (int, error) {
	response, err := http.Get(APICount)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return 0, errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	result := make(map[string]interface{})
	err = json.Unmarshal(data, &result)
	fmt.Printf("Count: %d\n", int(result["msg"].(float64)))
	return int(result["msg"].(float64)), nil
}

func requestUpdate(pID string, name string, age int, gender string) error {
	jsonData := map[string]interface{}{"user_id": pID, "name": name, "age": age, "gender": gender}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(APIUpdate, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	result := make(map[string]string)
	err = json.Unmarshal(data, &result)
	println("Update: " + result["msg"])
	return nil
}

func requestDelete(pID string) error {
	response, err := http.Get(APIDelete + "/" + pID)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return errors.New("Request fail")
	}
	data, _ := ioutil.ReadAll(response.Body)
	result := make(map[string]string)
	err = json.Unmarshal(data, &result)
	println("Delete: " + result["msg"])
	return nil
}

func requestDeleteAll() error {
	_, err := http.Get(APIDeleteAll)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return errors.New("Request fail")
	}
	fmt.Printf("Delete all persons success!\n")
	return nil
}
