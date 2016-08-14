package features

import(
	. "github.com/gucumber/gucumber"
)

func init() {
	initialNumber, results := 0,0

	Given(`^i have (\d+) cukes$`, func(cukes int) {
		initialNumber = cukes
	})

	When(`^I use the doubler$`, func() {
		results = initialNumber * 2
	})

	Then(`^I should have (\d+) cukes$`, func(i1 int) {
		T.Skip()
	})

}
