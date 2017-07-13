package storage

import "github.com/stretchr/testify/mock"

type TestStore struct {
	mock.Mock
}

func (s *TestStore) InitBucket(bucket string) error {
	args := s.Called(bucket)

	return args.Error(0)
}

func (s *TestStore) Get(bucket string, key string) (string, error) {
	args := s.Called(bucket, key)

	return args.String(0), args.Error(1)
}

func (s *TestStore) Put(bucket string, key string, value string) error {
	args := s.Called(bucket, key, value)

	return args.Error(0)
}

func (s *TestStore) GetAll(bucket string) (map[string]string, error) {
	args := s.Called(bucket)

	return args.Get(0).(map[string]string), args.Error(1)
}

func (s *TestStore) DeleteAll(bucket string) error {
	args := s.Called(bucket)

	return args.Error(0)
}
