package config

import "github.com/mike182uk/snpt/internal/platform/storage"

const (
	// TokenKey is the key of the github access token config value
	TokenKey = "gh-access-token"

	// LastSyncKey is the key of the last sync timestamp config value
	LastSyncKey = "last-sync"

	bucketName = "Config"
)

// Config stores and retrieves config values from the data store
type Config struct {
	storage storage.BucketKeyValueStore
}

// New returns a new Config
func New(storage storage.BucketKeyValueStore) (*Config, error) {
	config := Config{
		storage,
	}

	err := storage.InitBucket(bucketName)

	return &config, err
}

// Get retieves a config value
func (c *Config) Get(id string) (string, error) {
	val, err := c.storage.Get(bucketName, id)

	if err != nil {
		return "", err
	}

	return val, nil
}

// Set sets a config value
func (c *Config) Set(k string, v string) error {
	return c.storage.Put(bucketName, k, v)
}
