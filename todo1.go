
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

	i := Todo {1,2,3,4,5,6,7,8,9,10}
	for _, n: range Id {
		if Id == i{
			return true
		}else {
			fmt.Println("it doesn't exist")
		}
	}
	
}
