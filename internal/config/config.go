package config

import "github.com/mike182uk/snpt/internal/platform/storage"

// TokenKey is the key of the github access token config value
const TokenKey = "gh-access-token"

// LastSyncKey is the key of the last sync timestamp config value
const LastSyncKey = "last-sync"

// Config stores and retrieves config values from the data store
type Config struct {
	storage storage.BucketKeyValueStore
}

const bucketName = "Config"

// New returns a new Config instance
func New(storage storage.BucketKeyValueStore) (*Config, error) {
	config := &Config{
		storage: storage,
	}

	err := storage.InitBucket(bucketName)

	return config, err
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
func (c *Config) Set(key string, val string) error {
	return c.storage.Put(bucketName, key, val)
}
