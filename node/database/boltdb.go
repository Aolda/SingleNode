package database

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

func Boltdb() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Start a writable transaction.
	err = db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte("func"))
		if err != nil {
			return err
		}

		file, err := os.Open("./compiler/src/add.js")
		if err != nil {
			return err
		}
		defer file.Close()

		data := make([]byte, 100)
		_, err = file.Read(data)
		if err != nil {
			return err
		}

		// Insert the file contents into the bucket.
		err = b.Put([]byte("add.js"), data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
