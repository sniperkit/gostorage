/*
 * gostorage
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gostorage_test

import (
	"errors"
	"testing"

	"github.com/devfacet/gostorage"
	"github.com/devfacet/gostorage/drivers/local"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("should return a new storage", t, func() {
		s, err := gostorage.New(gostorage.Options{Driver: "local"})
		So(err, ShouldBeNil)
		So(s, ShouldNotBeNil)
	})

	Convey("should fail to return a new storage due to invalid driver", t, func() {
		s, err := gostorage.New(gostorage.Options{})
		So(err, ShouldBeError, errors.New("invalid storage driver "))
		So(s, ShouldBeNil)

		s, err = gostorage.New(gostorage.Options{Driver: "foo"})
		So(err, ShouldBeError, errors.New("invalid storage driver foo"))
		So(s, ShouldBeNil)
	})
}

func TestRegister(t *testing.T) {
	Convey("should register a new driver", t, func() {
		err := gostorage.Register("test", &local.Driver{})
		So(err, ShouldBeNil)
	})

	Convey("should fail to register a new driver", t, func() {
		err := gostorage.Register("", nil)
		So(err, ShouldBeError, errors.New("invalid driver name"))

		err = gostorage.Register("foo", nil)
		So(err, ShouldBeError, errors.New("invalid driver"))

		err = gostorage.Register("local", &local.Driver{})
		So(err, ShouldBeError, errors.New("driver local is already registered"))
	})
}
