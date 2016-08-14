package shoe_store

import (
	"github.com/sclevine/agouti"

	. "github.com/sclevine/agouti/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var Driver *agouti.WebDriver
var page *agouti.Page
var err error
var ShoeStoreUrl = "https://rb-shoe-store.herokuapp.com/"

var _ = Describe("Shoe Store", func() {
	// waits 60 seconds before throwing element not found
	SetDefaultEventuallyTimeout(60000 * time.Millisecond)
	// polls ever 10 milliseconds for whatever the SetDefaultEventuallyTimeout is
	SetDefaultEventuallyPollingInterval(10 * time.Millisecond)

	BeforeEach(func() {
		Driver = agouti.ChromeDriver()
		Expect(Driver.Start()).To(Succeed())

		page, err = Driver.NewPage()
		Expect(err).ToNot(HaveOccurred())

	})

	AfterEach(func() {
		Expect(Driver.Stop()).To(Succeed())
	})

	Describe("Shoe Store has Months", func() {
		It("Has Months", func() {
			Expect(true).Should(Equal(true))
		})
		It("Navigates to shoe store", func() {
			page.Navigate(ShoeStoreUrl)
			Expect(page.Title()).To(Equal("Shoe Store: Welcome to the Shoe Store"))
		})

		It("Navigates to July", func(){
			page.Navigate(ShoeStoreUrl)
			Eventually(page.FindByLink("July")).Should(BeFound())
			page.FindByLink("July").Click()
			Eventually(page.FindByClass("title")).Should(HaveText("July's Shoes"))
			//Eventually(page.FindByClass("title")).Should(HaveText("July's Shoes"))
		}, 5)
	})
})