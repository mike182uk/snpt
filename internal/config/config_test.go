package config

import (
	"errors"
	"testing"

	"github.com/mike182uk/snpt/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}

	storage.On("InitBucket", "Config").Return(nil)

	config, err := New(storage)

	storage.AssertExpectations(t)

	assert.IsType(t, &Config{}, config)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	config := &Config{storage: storage}
	key := "foo"
	value := "bar"

	storage.On("Get", "Config", key).Return(value, nil)

	result, err := config.Get(key)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, value, result)
}

func TestGetStorageErr(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	config := &Config{storage: storage}
	key := "foo"
	getErr := errors.New("Get Error")

	storage.On("Get", "Config", key).Return("", getErr)

	_, err := config.Get(key)

	storage.AssertExpectations(t)

	assert.Equal(t, getErr, err)
}

func TestSet(t *testing.T) {
	storage := &mocks.BucketKeyValueStore{}
	config := &Config{storage: storage}
	key := "foo"
	value := "bar"

	storage.On("Put", "Config", key, value).Return(nil)

	err := config.Set(key, value)

	storage.AssertExpectations(t)

	assert.Nil(t, err)
}
