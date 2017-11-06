/*
 * gostorage
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package local implements a local storage driver
package local

import (
	"fmt"
	"io"
	"os"
	gopath "path"

	"github.com/devfacet/gostorage"
)

func init() {
	gostorage.Register("local", &Driver{})
}

// Driver represents the storage driver
type Driver struct {
}

// NewWriter returns a new writer for the given path
func (d *Driver) NewWriter(path string) io.WriteCloser {
	// Init pipe
	pr, pw := io.Pipe()

	// Init writer
	w := &writer{path: path, perm: 0644, pw: pw, done: make(chan struct{})}

	// Copy
	go func() {
		defer close(w.done)

		// Create file
		base := gopath.Base(w.path)
		dir := gopath.Dir(w.path)
		fp := fmt.Sprintf("%s/%s", dir, base)

		if err := os.MkdirAll(dir, 0755); err != nil {
			pr.CloseWithError(fmt.Errorf("failed to create %s due to %s", dir, err.Error()))
		}

		f, err := os.Create(fp)
		if err != nil {
			pr.CloseWithError(fmt.Errorf("failed to create %s due to %s", fp, err.Error()))
		}
		defer f.Close()

		// Copy data
		_, err = io.Copy(f, pr)
		pr.CloseWithError(err)
	}()

	return w
}

type writer struct {
	path string
	perm os.FileMode
	pw   *io.PipeWriter
	done chan struct{}
}

// Write writes bytes from data
func (w *writer) Write(data []byte) (int, error) {
	return w.pw.Write(data)
}

// Close closes the writer
func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	<-w.done
	return nil
}
