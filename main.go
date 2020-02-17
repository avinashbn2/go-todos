package main

import (
	"flag"
	"fmt"
	"github.com/avina-sh/go-todos/models"
)
// Global slice to store todos
var todos = []models.Todo{}

// setup commands here
func InitCommands() {

	todoStr := flag.String("todo", "New TODO", "Enter Todo description")
	flag.Parse()
	fmt.Println(*todoStr)
	todo := models.Todo{*todoStr, false}
	todos = append(todos, todo)
	fmt.Println(todos)
}

func main() {
	InitCommands()

}