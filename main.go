package main

import (
	"flag"
	"fmt"

	"github.com/avina-sh/go-todos/models"
	"github.com/boltdb/bolt"
)

// Global slice to store todos
var todos = []models.Todo{}

// SFlag : Flag type to indicate if its value is set, to distinguish different flags
type SFlag struct {
	set   bool
	value string
}

// Set : Set value for the flag, (implments Flag interface)
func (s *SFlag) Set(x string) error {
	fmt.Println(x)
	s.set = true
	s.value = x
	return nil
}
func (s *SFlag) String() string {
	return s.value
}

var add SFlag

// InitCommands : setup commands here
func InitCommands(db *bolt.DB) {

	flag.Var(&add, "add", "Add new Todo")
	// flag.Int(&done, 1,  "Mark todo as done" )
	show := flag.Bool("show", false, "To Display Todos")
	flag.Parse()
	if add.set {
		todo := models.Todo{Name: add.value, Done: false}
		addTodo(todo, db)
	}
	if *show {
		models.DisplayTodos(db)

	}

}
func addTodo(todo models.Todo, db *bolt.DB) {
	err := todo.Save(db)
	if err != nil {
		fmt.Println(err)
	}
	todos = append(todos, todo)

}
func setUpDB() (*bolt.DB, error) {
	db, err := bolt.Open("todos.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Todos"))
		if err != nil {
			return fmt.Errorf("failed to create bucket %v", err)

		}
		return nil
	})

	return db, nil
}
func main() {
	db, err := setUpDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	InitCommands(db)

}
