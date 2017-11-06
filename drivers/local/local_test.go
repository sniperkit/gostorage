/*
 * gostorage
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package local_test

import (
	"io"
	"os"
	"testing"

	"github.com/devfacet/gostorage/drivers/local"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewWriter(t *testing.T) {
	Convey("should return a new writer", t, func() {
		f, err := os.Open("test.txt")
		So(err, ShouldBeNil)
		defer f.Close()

		d := local.Driver{}
		w := d.NewWriter("./tmp/test.txt")
		defer w.Close()
		defer os.RemoveAll("./tmp/")

		i, err := io.Copy(w, f)
		So(err, ShouldBeNil)
		So(i, ShouldBeGreaterThan, 0)
	})
}
