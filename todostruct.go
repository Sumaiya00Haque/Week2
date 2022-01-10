package main

import (
	"flag"
	"fmt"
	"time"
)

type ParsedTodoItem struct {
	Todo string
	date time.Time
	Done flag.Flag
}

var todoitem = flag.String("add", "homework", "Something worth doing")

func PrintItem(item *string) {
	fmt.Printf("You entered:\n\t%s", *todoitem)

}

func main() {
	fmt.Println("Welcome to ToDo list app")
	flag.Parse()

	PrintItem(todoitem)

}
