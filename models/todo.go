package models

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

type Todo struct {
	Name string
	Done bool

}
func DisplayTodos(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Todos"))

		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	return err
}
func (todo *Todo) Save(db *bolt.DB) error {
	fmt.Println("SAVE")

	// Store the user model in the user bucket using the username as the key.
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Todos"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		return b.Put([]byte(todo.Name), encoded)
	})
	return err
}