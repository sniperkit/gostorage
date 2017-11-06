/*
 * gostorage
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package gostorage is a library that provides a generic interface for various object storages
package gostorage

import (
	"fmt"
	"io"
)

// Options represents the options than can be set when creating a new storage
type Options struct {
	// Driver is the name of the storage driver
	Driver string
}

// New returns a new storage by the given options
func New(o Options) (*Storage, error) {
	// For concurrency
	driversMu.RLock()
	defer driversMu.RUnlock()

	// Init storage
	storage := Storage{
		isInit: true,
		driver: drivers[o.Driver],
	}

	if storage.driver == nil {
		return nil, fmt.Errorf("invalid storage driver %s", o.Driver)
	}

	return &storage, nil
}

// Storage represents a storage
type Storage struct {
	isInit bool
	driver Driver
}

// NewWriter returns a new writer for the given path
func (storage *Storage) NewWriter(path string) io.WriteCloser {
	return storage.driver.NewWriter(path)
}
