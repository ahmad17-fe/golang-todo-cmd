package view

import (
	"fmt"
	"golang-todo/controller"
	"golang-todo/model"
	"golang-todo/utils"
	"strconv"
)

func TodoApp() {
	var isLooping bool = true
	fmt.Println("---- Todo List ----")
	for isLooping {
		fmt.Println("Pilih Menu: ")
		fmt.Println("1. Menampilkan list")
		fmt.Println("2. Menambah list")
		fmt.Println("3. Mengubah list menjadi selesai")
		fmt.Println("4. Menghapus list")
		fmt.Println("5. Keluar")
		var option string = utils.Input("Pilih menu")
		fmt.Println("")
		switch option {
		case "1":
			listView()
		case "2":
			addView()
		case "3":
			updateView()
		case "4":
			removeView()
		case "5":
			isLooping = false
		default:
			fmt.Println("Menu tidak valid!")
		}
		fmt.Println("")
	}
}

func listView() {
	fmt.Println("Lists: ")
	res := controller.GetTodos()
	todos, ok := res.Data.([]model.Todo)

	if res.Status == utils.STATUS_SUCCESS {
		if ok {
			for _, todo := range todos {
				doneStatus := ""
				if todo.IsComplete {
					doneStatus = " | DONE"
				}
				fmt.Printf("%d. %s%s\n", todo.ID, todo.Title, doneStatus)
			}
		} else {
			fmt.Println("Failed to retrieve data.")
		}
	} else {
		fmt.Println(res.Message)
	}
}

func addView() {
	task := utils.Input("Masukkan task baru")

	var response = controller.CreateTodo(task)
	fmt.Println(response.Message)
}

func updateView() {
	idStr := utils.Input("Masukkan ID")
	idInt, _ := strconv.Atoi(idStr)
	respone := controller.UpdateTodo(idInt, model.Todo{
		IsComplete: true,
	})
	fmt.Println(respone.Message)
}

func removeView() {
	idStr := utils.Input("Masukkan ID")
	idInt, _ := strconv.Atoi(idStr)
	respone := controller.RemoveTodo(idInt)
	fmt.Println(respone.Message)
}
