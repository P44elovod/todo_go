package main

import "time"

//Todo is a model of Todo thing
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos is new type of data as a slice of todos
type Todos []Todo
