package storage

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// BoltStore is a Bolt backed bucket key value store
type BoltStore struct {
	boltDB *bolt.DB
	path   string
}

// NewBoltStore returns a new BoltStore instance
func NewBoltStore(path string) (*BoltStore, error) {
	boltDB, err := bolt.Open(path, 0644, nil)

	if err != nil {
		return nil, err
	}

	bs := &BoltStore{
		boltDB: boltDB,
		path:   path,
	}

	return bs, nil
}

// InitBucket initialises a bucket
func (bs *BoltStore) InitBucket(bucket string) error {
	return bs.boltDB.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))

		return err
	})
}

// Get retrieves a value from a bucket
func (bs *BoltStore) Get(bucket string, key string) (val string, err error) {
	err = bs.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return fmt.Errorf("%s is not a bucket", bucket)
		}

		val = string(b.Get([]byte(key)))

		return nil
	})

	return
}

// GetAll retrieves all key-value pairs from a bucket
func (bs *BoltStore) GetAll(bucket string) (vals map[string]string, err error) {
	vals = make(map[string]string)

	err = bs.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			vals[string(k)] = string(v)
		}

		return nil
	})

	return
}

// Put stores a key-value pair in a bucket
func (bs *BoltStore) Put(bucket string, key string, value string) error {
	return bs.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return fmt.Errorf("%s is not a bucket", bucket)
		}

		return b.Put([]byte(key), []byte(value))
	})
}

// DeleteAll deletes all key-value pairs stored in a bucket
func (bs *BoltStore) DeleteAll(bucket string) error {
	return bs.boltDB.Batch(func(tx *bolt.Tx) error {
		if err := tx.DeleteBucket([]byte(bucket)); err != nil {
			return err
		}

		if _, err := tx.CreateBucket([]byte(bucket)); err != nil {
			return err
		}

		return nil
	})
}
