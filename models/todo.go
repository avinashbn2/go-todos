package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

type Todo struct {
	Id   int
	Name string
	Done bool
}

// Display Todos from db
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

// Save : function to save todoitem to DB
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
		id, _ := b.NextSequence()
		todo.Id = int(id)
		return b.Put(itob(todo.Id), encoded)
	})
	return err
}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
