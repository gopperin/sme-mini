package bolt

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Bblot Bblot
var Bblot Database

// Database is a persist-level wrapper for the bolt database, providing
// extra information such as a version number.
type Database struct {
	*bolt.DB
}

// GetData GetData
func (db *Database) GetData(bucket, key string) string {
	var ret string
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if nil == bucket {
			fmt.Println("bucket is nil")
			ret = ""
			return nil
		}
		bytes := bucket.Get([]byte(key))
		if nil == bytes {
			fmt.Println("key is nil")
			ret = ""
		} else {
			ret = string(bytes)
		}

		return nil
	})
	return ret
}

// GetAllKey GetAllKey
func (db *Database) GetAllKey(bucket string) string {
	var ret string
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if nil == bucket {
			fmt.Println("bucket is nil")
			return nil
		}
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("key = %s,value = %s\n", k, v)
		}
		return nil
	})
	return ret
}

// PutData PutData
func (db *Database) PutData(bucket, key, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Close closes the database.
func (db *Database) Close() error {
	return db.DB.Close()
}

// OpenDatabase opens a database and validates its metadata.
func OpenDatabase(filename string) (*Database, error) {
	// Open the database using a 3 second timeout (without the timeout,
	// database will potentially hang indefinitely.
	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}

	// Check the metadata.
	boltDB := &Database{
		DB: db,
	}

	Bblot = *boltDB

	return boltDB, nil
}
