package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSquare(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an integer value as input", t, func() {
		x := 3

		Convey("When call func Square", func() {
			value := Square(x)

			Convey("Then output should be input square", func() {
				So(value, ShouldEqual, 9)
			})
		})
	})
}