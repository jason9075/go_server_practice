package model

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type Person struct {
	ID     string
	Name   string `json:"name, string"`
	Age    int    `json:"age, int"`
	Gender string `json:"gender, string"`
}

func InsertPerson(p Person) string {
	db := GetDb()
	pid, _ := uuid.NewV4()
	p.ID = pid.String()
	mu.Lock()
	defer mu.Unlock()
	db.personMap[pid.String()] = p
	return pid.String()
}

func QueryPerson(pid string) Person {
	db := GetDb()
	return db.personMap[pid]
}

func UpdatePerson(userID string, payload map[string]interface{}) bool {
	db := GetDb()
	oldPerson, ok := db.personMap[userID]
	if ok {
		newName := checkNil(payload["name"], oldPerson.Name).(string)
		newGender := checkNil(payload["gender"], oldPerson.Gender).(string)
		newAge := int(checkNil(payload["age"], oldPerson.Age).(float64))
		newPerson := Person{ID: userID, Name: newName, Age: newAge, Gender: newGender}
		db.personMap[userID] = newPerson
		return true
	}
	return false
}

func DeletePerson(pid string) bool {
	db := GetDb()
	_, ok := db.personMap[pid]
	if ok {
		delete(db.personMap, pid)
		return true
	}
	return false
}

func ReadAllPerson() []Person {
	db := GetDb()
	persons := make([]Person, 0, len(db.personMap))
	for _, person := range db.personMap {
		persons = append(persons, person)
	}
	return persons
}

func DeleteAllPerson() {
	db := GetDb()
	db.personMap = make(map[string]Person, 1000)
}

func PrintAllPerson() {
	db := GetDb()
	fmt.Println("Current Persons:")
	for _, person := range db.personMap {
		fmt.Println(person)
	}
}
