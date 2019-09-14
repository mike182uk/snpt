package snippet

import (
	"errors"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mike182uk/snpt/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	snpt = Snippet{
		Id:          "foo",
		Filename:    "bar",
		Description: "baz",
		Content:     "qux",
	}
	snptSerialized, _ = proto.Marshal(&snpt)
	snpts             = Snippets{snpt}
	snptsMap          = map[string]string{
		snpt.Id: string(snptSerialized),
	}
)

func TestNewStore(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}

	storage.On("InitBucket", "Snippets").Return(nil)

	expected := &Store{}
	result, err := NewStore(storage)

	storage.AssertExpectations(t)

	assert.IsType(t, expected, result)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}
	key := "foo"

	storage.On("Get", "Snippets", key).Return(string(snptSerialized), nil)

	result, err := store.Get(key)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
	assert.IsType(t, Snippet{}, result)
	assert.Equal(t, snpt.GetId(), result.GetId())
}

func TestGetStorageErr(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}
	key := "foo"
	getErr := errors.New("Get Error")

	storage.On("Get", "Snippets", key).Return("", getErr)

	_, err := store.Get(key)

	storage.AssertExpectations(t)

	assert.Equal(t, getErr, err)
}

func TestGetAll(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}

	storage.On("GetAll", "Snippets").Return(snptsMap, nil)

	result, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, snpts, result)
}

func TestGetAllStorageErr(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}
	getAllErr := errors.New("GetAll Error")

	storage.On("GetAll", "Snippets").Return(map[string]string{}, getAllErr)

	_, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Equal(t, getAllErr, err)
}

func TestGetAllUnserializeErr(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}
	invalidSnptsMap := map[string]string{
		"foo": "bar",
	}

	storage.On("GetAll", "Snippets").Return(invalidSnptsMap, nil)

	_, err := store.GetAll()

	storage.AssertExpectations(t)

	assert.Error(t, err)
}

func TestPut(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}

	storage.On("Put", "Snippets", snpt.GetId(), string(snptSerialized)).Return(nil)

	err := store.Put(snpt)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestDeleteAll(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	store := &Store{storage: storage}

	storage.On("DeleteAll", "Snippets").Return(nil)

	err := store.DeleteAll()

	storage.AssertExpectations(t)

	assert.Nil(t, err)
}
