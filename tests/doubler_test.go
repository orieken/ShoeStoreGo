package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
)

func TestDoubler(t *testing.T) {
	initialValue, result := 0, 0

	Convey("Given I have an initial value of 3", t, func() {
		initialValue = 3
		Convey("When I run it thru the doubler", func() {
			result = initialValue * 2
			So(result, should.Equal, initialValue * 2)
		})
	})
}