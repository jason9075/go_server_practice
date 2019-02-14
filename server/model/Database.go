package model

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type database struct {
	personMap map[string]Person
}

var instance *database
var mu sync.Mutex
var initialized uint32

func GetDb() *database {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &database{make(map[string]Person, 1000)}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func (db *database) ReadAll() []Person {
	persons := make([]Person, 0, len(db.personMap))
	for _, person := range db.personMap {
		persons = append(persons, person)
	}
	return persons
}

func (db *database) DeleteAll() {
	db.personMap = make(map[string]Person, 1000)
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
