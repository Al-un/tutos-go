package main

import "fmt"

var currentId int
var todos Todos

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)

	return t
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	// empty Todo if not found
	return Todo{}
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find id %d to delete from database", id)
}

func init() {
	RepoCreateTodo(Todo{Name: "Presentation"})
	RepoCreateTodo(Todo{Name: "Meetup"})
}
