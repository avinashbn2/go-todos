package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

// Todo struct for storing todo data
type Todo struct {
	ID   int
	Name string
	Done bool
}

// DisplayTodos from db
func DisplayTodos(db *bolt.DB, done bool) error {

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Todos"))
		// todos := []*Todo{}
		b.ForEach(func(k, v []byte) error {
			temp := &Todo{}
			json.Unmarshal(v, temp)
			if done {
				if temp.Done {
					fmt.Println(string(v))
				}
			} else {
				fmt.Println(string(v))

			}

			return nil
		})
		return nil
	})
	return err
}

// Done : function to update a Todo as done
func Done(db *bolt.DB, done int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Todos"))
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		v := b.Get(itob(done))
		todo := &Todo{}
		err = json.Unmarshal([]byte(v), todo)
		if err != nil {
			fmt.Println("Invalid id")
			return err
		}
		todo.Done = true
		encoded, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		return b.Put(itob(todo.ID), encoded)

	})

	return err
}

// Save : function to save todoitem to DB
func (todo *Todo) Save(db *bolt.DB) error {

	// Store the user model in the user bucket using the username as the key.
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Todos"))
		if err != nil {
			return err
		}
		bT := tx.Bucket([]byte("Todos"))
		id, _ := bT.NextSequence()

		todo.ID = int(id)

		encoded, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		return bT.Put(itob(todo.ID), encoded)
	})
	return err
}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
