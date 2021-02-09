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
	var _ret string
	db.View(func(tx *bolt.Tx) error {
		_bucket := tx.Bucket([]byte(bucket))
		if nil == _bucket {
			fmt.Println("bucket is nil")
			_ret = ""
			return nil
		}
		_bytes := _bucket.Get([]byte(key))
		if nil == _bytes {
			fmt.Println("key is nil")
			_ret = ""
		} else {
			_ret = string(_bytes)
		}

		return nil
	})
	return _ret
}

// GetAllKey GetAllKey
func (db *Database) GetAllKey(bucket string) string {
	var _ret string
	db.View(func(tx *bolt.Tx) error {
		_bucket := tx.Bucket([]byte(bucket))
		if nil == _bucket {
			fmt.Println("bucket is nil")
			return nil
		}
		_cursor := _bucket.Cursor()
		for k, v := _cursor.First(); k != nil; k, v = _cursor.Next() {
			fmt.Printf("key = %s,value = %s\n", k, v)
		}
		return nil
	})
	return _ret
}

// PutData PutData
func (db *Database) PutData(bucket, key, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		err = _bucket.Put([]byte(key), []byte(value))
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
