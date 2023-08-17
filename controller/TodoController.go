package controller

import (
	"golang-todo/model"
	"golang-todo/utils"
)

func GetTodos() *utils.ResponseCallback {
	var todos = model.GetTodo()
	if len(todos) > 0 {
		response := utils.NewResponseCallback(&utils.ResponseCallback{
			Status:     utils.STATUS_SUCCESS,
			StatusCode: 200,
			Data:       todos,
			Message:    "Success get the data",
		})
		return response
	}
	response := utils.NewResponseCallback(&utils.ResponseCallback{
		Status:     utils.STATUS_ERROR,
		StatusCode: 404,
		Message:    "No task found!",
	})
	return response
}

func CreateTodo(todo string) *utils.ResponseCallback {
	var response *utils.ResponseCallback

	if todo == "" {
		callback := &utils.ResponseCallback{
			Status:     utils.STATUS_ERROR,
			StatusCode: 406,
			Message:    "Failed to add new list!",
		}
		response = utils.NewResponseCallback(callback)
		return response
	}
	model.AddTodo(todo)
	response = utils.NewResponseCallback(&utils.ResponseCallback{
		Status:     utils.STATUS_SUCCESS,
		StatusCode: 200,
		Message:    "Successfully added a new list",
	})
	return response
}

func UpdateTodo(ID int, value model.Todo) *utils.ResponseCallback {
	var response *utils.ResponseCallback
	todo, err := model.GetTodoById(ID)
	if err != nil {
		todo.ID = ID
		todo.IsComplete = value.IsComplete
		model.UpdateTodo(todo)

		callback := &utils.ResponseCallback{
			Status:     utils.STATUS_SUCCESS,
			StatusCode: 200,
			Message:    "Success update task!",
		}
		response = utils.NewResponseCallback(callback)
	} else {
		callback := &utils.ResponseCallback{
			Status:     utils.STATUS_ERROR,
			StatusCode: 404,
			Message:    "Task not found!",
		}
		response = utils.NewResponseCallback(callback)
	}
	return response
}

func RemoveTodo(ID int) *utils.ResponseCallback {
	var response *utils.ResponseCallback
	_, err := model.GetTodoById(ID)
	if err != nil {
		model.DeleteTodo(ID)

		callback := &utils.ResponseCallback{
			Status:     utils.STATUS_SUCCESS,
			StatusCode: 200,
			Message:    "Success delete task!",
		}
		response = utils.NewResponseCallback(callback)
	} else {
		callback := &utils.ResponseCallback{
			Status:     utils.STATUS_ERROR,
			StatusCode: 404,
			Message:    "Task not found!",
		}
		response = utils.NewResponseCallback(callback)
	}
	return response
}
