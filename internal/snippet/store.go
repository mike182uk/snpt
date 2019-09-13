package snippet

import (
	"github.com/golang/protobuf/proto"

	"github.com/mike182uk/snpt/internal/platform/storage"
)

const bucketName = "Snippets"

// Store retrieves, creates and deletes snippets stored in the data store
type Store struct {
	storage storage.BucketKeyValueStore
}

// NewStore returns a new Store instance
func NewStore(storage storage.BucketKeyValueStore) (*Store, error) {
	store := &Store{
		storage: storage,
	}

	err := storage.InitBucket(bucketName)

	return store, err
}

// Get retrieves a snippet from the store
func (s *Store) Get(id string) (snpt Snippet, err error) {
	b, err := s.storage.Get(bucketName, id)

	if err != nil {
		return
	}

	err = proto.Unmarshal([]byte(b), &snpt)

	return
}

// GetAll retrieves all snippets from the store
func (s *Store) GetAll() (snpts Snippets, err error) {
	bs, err := s.storage.GetAll(bucketName)

	if err != nil {
		return
	}

	for _, b := range bs {
		var snpt Snippet

		err = proto.Unmarshal([]byte(b), &snpt)

		if err != nil {
			return
		}

		snpts = append(snpts, snpt)
	}

	return
}

// Put stores a snippet in the store
func (s *Store) Put(snpt Snippet) error {
	b, err := proto.Marshal(&snpt)

	if err != nil {
		return err
	}

	return s.storage.Put(bucketName, snpt.GetId(), string(b))
}

// DeleteAll deletes all snippets in the store
func (s *Store) DeleteAll() error {
	return s.storage.DeleteAll(bucketName)
}
