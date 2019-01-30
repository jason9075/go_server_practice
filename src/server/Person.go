package main

type Person struct {
	ID     string
	Name   string `json:"name, string"`
	Age    int    `json:"age, int"`
	Gender string `json:"gender, string"`
}
