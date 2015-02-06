package demo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	Convey("Given integers not divisible by 3 or 5", t, func() {
		Convey("It should return the number", func() {
			So(fizzbuzz(1), ShouldEqual, "1")
			So(fizzbuzz(2), ShouldEqual, "2")
			So(fizzbuzz(4), ShouldEqual, "4")
		})
	})

	Convey("For multiples of 3", t, func() {
		Convey("It should return the \"Fizz\"", func() {
			So(fizzbuzz(3), ShouldEqual, "Fizz")
			So(fizzbuzz(6), ShouldEqual, "Fizz")
			So(fizzbuzz(9), ShouldEqual, "Fizz")
			So(fizzbuzz(99), ShouldEqual, "Fizz")
		})
	})

	Convey("For multiples of 5", t, func() {
		Convey("It should return \"buzz\"", func() {
			So(fizzbuzz(5), ShouldEqual, "Buzz")
			So(fizzbuzz(10), ShouldEqual, "Buzz")
			So(fizzbuzz(20), ShouldEqual, "Buzz")
		})
	})

	Convey("For multiples of both 3 and 5", t, func() {
		Convey("It should return \"fizzbuzz\"", func() {
			So(fizzbuzz(15), ShouldEqual, "FizzBuzz")
		})
	})
}
