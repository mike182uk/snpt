package storage

// BucketKeyValueStore represents a bucket key value store
type BucketKeyValueStore interface {
	InitBucket(bucket string) error
	Get(bucket string, key string) (string, error)
	Put(bucket string, key string, value string) error
	GetAll(bucket string) (map[string]string, error)
	DeleteAll(bucket string) error
}
