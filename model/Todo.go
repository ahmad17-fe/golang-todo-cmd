package model

import "fmt"

type Todo struct {
	ID         int
	Title      string
	IsComplete bool
}

var Todos []Todo

func GetTodo() []Todo {
	return Todos
}

func GetTodoById(ID int) (Todo, error) {
	for _, todo := range Todos {
		if todo.ID == ID {
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("Todo with ID %d not found", ID)
}

func AddTodo(todo string) {
	var lastId int
	if len(Todos) > 0 {
		lastId = Todos[len(Todos)-1].ID + 1
	} else {
		lastId = 1
	}

	var body = Todo{
		ID:         lastId,
		Title:      todo,
		IsComplete: false,
	}

	Todos = append(Todos, body)
}

func UpdateTodo(data Todo) {
	for i, todo := range Todos {
		if todo.ID == data.ID {
			Todos[i] = data
		}
	}
}

func DeleteTodo(ID int) {
	var getIndex int = -1
	for i, todo := range Todos {
		if todo.ID == ID {
			getIndex = i
		}
	}
	if getIndex != -1 {
		Todos = append(Todos[:getIndex], Todos[getIndex+1:]...)
	}
}
