package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	bolt "go.etcd.io/bbolt"
)

var testBucket = []byte("test")

func TestNewStore(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	store, err := NewBoltStore(dbPath)
	defer store.boltDB.Close() //nolint: staticcheck

	assert.Nil(t, err)
	assert.IsType(t, &BoltStore{}, store)
}

func TestNewStoreErr(t *testing.T) {
	_, err := NewBoltStore("")

	assert.NotNil(t, err)
}

func TestInitBucket(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	if err = store.InitBucket(string(testBucket)); err != nil {
		t.Fatal(err)
	}

	if err = store.boltDB.Close(); err != nil {
		t.Fatal(err)
	}

	db, err := getBoltDB(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(testBucket) // will return error if bucket does not exist
	})

	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	db, err := getBoltDB(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	key := "foo"
	value := "bar"

	createTestBucket(db)

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(testBucket)

		return b.Put([]byte(key), []byte(value))
	}); err != nil {
		t.Fatal(err)
	}

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer store.boltDB.Close()

	result, _ := store.Get(string(testBucket), key)

	assert.Equal(t, value, result)
}

func TestGetInvalidBucketErr(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	invalidBucket := "foo"

	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer store.boltDB.Close()

	_, err = store.Get(invalidBucket, "bar")

	assert.EqualError(t, err, fmt.Sprintf("%s is not a bucket", invalidBucket))
}

func TestGetAll(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	db, err := getBoltDB(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	createTestBucket(db)

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(testBucket)

		if err = b.Put([]byte("foo"), []byte("bar")); err != nil {
			return err
		}

		if err = b.Put([]byte("baz"), []byte("qux")); err != nil {
			return err
		}

		return err
	}); err != nil {
		t.Fatal(err)
	}

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer store.boltDB.Close()

	expected := map[string]string{"foo": "bar", "baz": "qux"}
	result, err := store.GetAll(string(testBucket))

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestPut(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	key := "foo"
	value := "bar"

	// create test bucket
	db, err := getBoltDB(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	createTestBucket(db)

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	// store item in bucket
	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	if err = store.Put(string(testBucket), key, value); err != nil {
		t.Fatal(err)
	}

	if err = store.boltDB.Close(); err != nil {
		t.Fatal(err)
	}

	// retrieve item from bucket
	db, err = getBoltDB(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(testBucket)

		result := b.Get([]byte(key))

		assert.Equal(t, value, string(result))

		return nil
	})
}

func TestDeleteAll(t *testing.T) {
	dbDir, dbPath := getDBPath()
	defer cleanupTestDB(dbDir)

	db, err := getBoltDB(dbPath)

	createTestBucket(db)

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(testBucket)

		if err = b.Put([]byte("foo"), []byte("bar")); err != nil {
			return err
		}

		if err = b.Put([]byte("baz"), []byte("qux")); err != nil {
			return err
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	store, err := NewBoltStore(dbPath)

	if err != nil {
		t.Fatal(err)
	}

	defer store.boltDB.Close()

	if err := store.DeleteAll(string(testBucket)); err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{}
	result, _ := store.GetAll(string(testBucket))

	assert.Equal(t, expected, result)
}

func getDBPath() (string, string) {
	dbDir, err := ioutil.TempDir("", "snpt-test")

	if err != nil {
		panic(err)
	}

	return dbDir, path.Join(dbDir, "snpt.db")
}

func getBoltDB(path string) (*bolt.DB, error) {
	return bolt.Open(path, 0644, nil)
}

func cleanupTestDB(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}

func createTestBucket(db *bolt.DB) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(testBucket)

		return err
	}); err != nil {
		panic(err)
	}
}
