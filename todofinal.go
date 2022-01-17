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
}

func PrintTodoList(todoList interface{}) {

	fmt.Println("Todolist:", todoList)
}

// add necessary todo := Todo{parameters
func AddTodo() {
	// write code to add todo in `todoList`` slice
	todo := Todo{
		Id:   1,
		Name: "sleep all day",
		Date: time.Now(),
		Done: false,
	}

	todoList := []interface{}{[]int{1, 2}, []string{"wakeup", "breakfast"}, []time.Time{time.Now()}, []bool{true, false}}
	todoList = append(todoList, todo)
	PrintTodoList(todoList)
	MarkAsDone(todo.Id)
}

// add necessary parameters
func MarkAsDone(t int) {
	// write code to change `Done` flag to true of a todo from todoList
	if t == 1 {
		fmt.Println("Done!!")
	} else {
		fmt.Println("It doesn't exist")
	}

}
