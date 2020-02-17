package main

import (
	"flag"
	"fmt"
	"github.com/avina-sh/go-todos/models"
	"github.com/boltdb/bolt"
)
// Global slice to store todos
var todos = []models.Todo{}

var DB *bolt.DB
// setup commands here
func InitCommands(db *bolt.DB) {

	todoStr := flag.String("todo", "New TODO", "Enter Todo description")
	flag.Parse()

	todo := models.Todo{*todoStr, false}
	err := todo.Save(db)
	if err!=nil {fmt.Println(err)}
	todos = append(todos, todo)
	models.DisplayTodos(db)
}
func setUpDB() (*bolt.DB, error){
	db, err := bolt.Open("todos.db", 0600, nil)
	if err!=nil {
		return nil, fmt.Errorf("Failed to connect to db %v", err)
	}
	err = db.Update( func(tx *bolt.Tx) error {
		_, err:= tx.CreateBucketIfNotExists([]byte("Todos"))
		if err!= nil{
			return fmt.Errorf("failed to create bucket %v", err)

		}
		return nil
	})

	return db, nil
}
func main() {
	db, err := setUpDB()
	if err!=nil {
		fmt.Println(err)
	}
	defer  db.Close()
	InitCommands(db)


}