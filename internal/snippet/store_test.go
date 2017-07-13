package snippet

import (
	"errors"
	"testing"

	"github.com/mike182uk/snpt/internal/platform/storage"
	"github.com/stretchr/testify/assert"
)

var (
	snpt = Snippet{
		ID:          "foo",
		Filename:    "bar",
		Description: "baz",
		Content:     "qux",
	}
	snptJSON  = `{"id":"foo","filename":"bar","description":"baz","content":"qux"}`
	snpts     = Snippets{snpt}
	snptsJSON = map[string]string{
		snpt.ID: snptJSON,
	}
)

func TestNewStore(t *testing.T) {
	storage := &storage.TestStore{}

	storage.On("InitBucket", "Snippets").Return(nil)

	expected := &Store{}
	result, err := NewStore(storage)

	storage.AssertExpectations(t)

	assert.IsType(t, expected, result)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}
	key := "foo"

	storage.On("Get", "Snippets", key).Return(snptJSON, nil)

	result, err := store.Get(key)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, snpt, result)
}

func TestGetStorageErr(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}
	key := "foo"
	getErr := errors.New("Get Error")

	storage.On("Get", "Snippets", key).Return("", getErr)

	_, err := store.Get(key)

	storage.AssertExpectations(t)

	assert.Equal(t, getErr, err)
}

func TestGetAll(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}

	storage.On("GetAll", "Snippets").Return(snptsJSON, nil)

	result, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, snpts, result)
}

func TestGetAllStorageErr(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}
	getAllErr := errors.New("GetAll Error")

	storage.On("GetAll", "Snippets").Return(map[string]string{}, getAllErr)

	_, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Equal(t, getAllErr, err)
}

func TestGetAllDecodingErr(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}
	invalidSnptsJSON := map[string]string{
		"foo": "bar",
	}

	storage.On("GetAll", "Snippets").Return(invalidSnptsJSON, nil)

	_, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Error(t, err)
}

func TestPut(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}

	storage.On("Put", "Snippets", snpt.ID, snptJSON).Return(nil)

	err := store.Put(snpt)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestDeleteAll(t *testing.T) {
	storage := &storage.TestStore{}
	store := &Store{storage: storage}

	storage.On("DeleteAll", "Snippets").Return(nil)

	err := store.DeleteAll()

	storage.AssertExpectations(t)

	assert.Nil(t, err)
}
