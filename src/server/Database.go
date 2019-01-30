package main

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type database struct {
	personMap map[string]Person
}

var instance *database

func GetDb() *database {
	if instance == nil {
		instance = &database{make(map[string]Person, 1000)}
	}
	return instance
}

func (db *database) Insert(p Person) string {
	pid, _ := uuid.NewV4()
	p.ID = pid.String()
	db.personMap[pid.String()] = p
	return pid.String()
}

func (db *database) Query(pid string) Person {
	return db.personMap[pid]
}

func (db *database) ReadAll() []Person {
	persons := make([]Person, 0, len(db.personMap))
	for _, person := range db.personMap {
		persons = append(persons, person)
	}
	return persons
}

func (db *database) Update(userID string, payload map[string]interface{}) bool {
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

func (db *database) Delete(pid string) bool {
	_, ok := db.personMap[pid]
	if ok {
		delete(db.personMap, pid)
		return true
	}
	return false
}

func (db *database) PrintAll() {
	fmt.Println("Current Persons:")
	for _, person := range db.personMap {
		fmt.Println(person)
	}
}

func checkNil(targetValue interface{}, defaultValue interface{}) interface{} {
	if targetValue == nil {
		return defaultValue
	}
	return targetValue
}
