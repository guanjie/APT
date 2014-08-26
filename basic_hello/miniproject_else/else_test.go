// else_test.go file for learning purpose.
package demo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestElse(t *testing.T) {

	Convey("For number that's greater than 0", t, func() {
		Convey("It should return 1", func() {
			So(comparedToZero(100), ShouldEqual, 1)
			So(comparedToZero(121212), ShouldEqual, 1)
			So(comparedToZero(123123123), ShouldEqual, 1)
		})
	})

	Convey("For number that's equal to 0", t, func() {
		Convey("It should return 0", func() {
			So(comparedToZero(0), ShouldEqual, 0)
		})
	})

	Convey("For number that's less than 0", t, func() {
		Convey("It should return -1", func() {
			So(comparedToZero(-100), ShouldEqual, -1)
			So(comparedToZero(-123), ShouldEqual, -1)
			So(comparedToZero(-1), ShouldEqual, -1)
		})
	})
}
