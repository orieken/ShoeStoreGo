package shoe_store

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestLandingPage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shoe Store Suite")
}

var _ = Describe("Shoe Store", func() {

	Describe("Landing Page has a title", func() {
		It("Has a Title", func() {
			Expect(true).Should(Equal(true))
		})
	})
})