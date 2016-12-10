package main

import (
	"strings"

	"github.com/boltdb/bolt"
)

type database struct {
	boltDB *bolt.DB
}

func newDatabase(boltDB *bolt.DB) *database {
	return &database{
		boltDB: boltDB,
	}
}

func (db *database) init(bucketNames []string) error {
	return db.boltDB.Batch(func(tx *bolt.Tx) error {
		var err error

		for _, v := range bucketNames {
			_, err = tx.CreateBucketIfNotExists([]byte(v))
		}

		return err
	})
}

func (db *database) get(k string) (string, error) {
	s := strings.Split(k, ".")

	bucket := s[0]
	key := s[1]

	tx, err := db.boltDB.Begin(false)

	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	b := tx.Bucket([]byte(bucket))
	v := b.Get([]byte(key))

	return string(v), nil
}

func (db *database) getAll(k string) (map[string]string, error) {
	vals := make(map[string]string)

	tx, err := db.boltDB.Begin(false)

	if err != nil {
		return vals, err
	}

	defer tx.Rollback()

	b := tx.Bucket([]byte(k))

	b.ForEach(func(k, v []byte) error {
		vals[string(k)] = string(v)

		return nil
	})

	return vals, nil
}

func (db *database) set(k string, v string) error {
	s := strings.Split(k, ".")

	err := db.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s[0]))
		err := b.Put([]byte(s[1]), []byte(v))

		return err
	})

	return err
}

func (db *database) empty(k string) error {
	return db.boltDB.Batch(func(tx *bolt.Tx) error {
		if err := tx.DeleteBucket([]byte(k)); err != nil {
			return err
		}

		if _, err := tx.CreateBucket([]byte(k)); err != nil {
			return err
		}

		return nil
	})
}
