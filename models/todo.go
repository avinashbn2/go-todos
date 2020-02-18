package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

type Todo struct {
	Id   int    `json:id`
	Name string `json:name`
	Done bool   `json:done`
}

// DisplayTodos from db
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
		return b.Put(itob(todo.Id), encoded)

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

		todo.Id = int(id)

		encoded, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		return bT.Put(itob(todo.Id), encoded)
	})
	return err
}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
