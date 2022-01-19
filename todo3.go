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

	PrintTodoList()
}

func PrintTodoList() {
	fmt.Println(todoList)
}

// add necessary parameters
func AddTodo(Id int, Name string, Date time.Time, Done bool) {
	// write code to add todo in `todoList`` slice
	todo := Todo{
		Id:   1,
		Name: "Wake Up",
		Date: time.Now(),
		Done: false,
	}

	todoList = append(todoList, todo)

}

// add necessary parameters
func MarkAsDone( Id int)bool {
	// write code to change `Done` flag to true of a todo from todoList

	
	for _, todo: range todolist {
		if Id == todo.Id{
			return true
		}else {
			fmt.Println("it doesn't exist")
		}
	}
	
}
