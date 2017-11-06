/*
 * gostorage
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gostorage

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Driver)
)

// Driver is the interface that must be implemented by storage dirvers
type Driver interface {
	// NewWriter returns a new writer for the given path
	NewWriter(path string) io.WriteCloser
}

// Register registers a storage driver
func Register(name string, driver Driver) error {
	// For concurrency
	driversMu.Lock()
	defer driversMu.Unlock()

	if name == "" {
		return errors.New("invalid driver name")
	} else if driver == nil {
		return errors.New("invalid driver")
	} else if _, ok := drivers[name]; ok {
		return fmt.Errorf("driver %s is already registered", name)
	}

	drivers[name] = driver

	return nil
}
