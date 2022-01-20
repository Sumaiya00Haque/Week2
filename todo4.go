package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Id   int
	Name string
	Date time.Time
	Done bool
}

var todoList []Todo = make([]Todo, 0)

func main() {
	// call other functions here
	AddTodo()
	MarkAsDone()
	PrintTodoList()
	AddTodo()
	PrintTodoList()
}

func PrintTodoList() {
	fmt.Println(todoList)
}

// add necessary parameters
func AddTodo(todo Todo) {
	// write code to add todo in `todoList`` slice
	// todo := Todo{
	// 	Id:   1,
	// 	Name: "Wake Up",
	// 	Date: time.Now(),
	// 	Done: false,
	// }

	todoList = append(todoList, todo)

}

// add necessary parameters
func MarkAsDone( Id int)bool {
	// write code to change `Done` flag to true of a todo from todoList

	
	for _, todo: range todolist {
		if Id == todo.Id{
			todo.Done = true
		}else {
			fmt.Println("it doesn't exist")
		}
	}
	
}
