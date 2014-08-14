//1_test.go file for testing purpose.

package erictest

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test1(t *testing.T) {
	Convey("For number 15", t, func() {
		Convey("It should give the number", func() {
			So(Increment(15), ShouldEqual, 20)
		})
	})
}

func Test2(t *testing.T) {
	Convey("For number 23", t, func() {
		Convey("It should give the number", func() {
			So(Increment(23), ShouldEqual, 28)
		})
	})
}
