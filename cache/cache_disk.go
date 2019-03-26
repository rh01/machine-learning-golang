package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("./embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// crreeate a bucket in the boltdb file for our data
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("mybucket"))
		if err != nil {
			return fmt.Errorf("create buckeet: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mybucket"))
		if err := b.Put([]byte("mykeye"), []byte("myvalue")); err != nil {
			return fmt.Errorf("put buckeet: %s", err)
		}
		return nil

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("mybucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Key: %s, Value: %s\n", string(k), string(v))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
