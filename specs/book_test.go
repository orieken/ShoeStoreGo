package book_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/types"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func HaveTitle(n string) types.GomegaMatcher {
	return WithTransform(func(book Book) string {
		return book.Title
	}, Equal(n))
}
func HaveAnAuthor(n string) types.GomegaMatcher {
	return WithTransform(func(book Book) string {
		return book.Author
	}, Equal(n))
}
func HavePageCount(ageMatcher types.GomegaMatcher) types.GomegaMatcher {
	// ageMatcher gives callers flexibility to use Equal, BeNumerically, etc.
	return WithTransform(func(book Book) int {
		return book.Pages
	}, ageMatcher)
}

func(book Book) printBookDetails() string {
	return fmt.Sprintf("Book Title: %v", book.Title)
}

var _ = Describe("Book", func() {
	var (
		shortBook Book
		longBook Book
		emptyBook Book
	)

	BeforeEach(func() {
		shortBook = Book{
			Title: "Book Title 1",
			Author: "John Doe",
			Pages: 123,
		}
		longBook = Book{
			Title: "Book Title 2",
			Author: "John Doe",
			Pages: 1023,
		}
		emptyBook = Book{
			Title: "Book Title 1",
			Author: "John Doe",
			Pages: 1,
		}
	})

	Describe("Book has a title", func() {
		It("Has a Title", func() {
			Expect(shortBook).Should(HaveTitle("Book Title 1"))
		})

		It("details are", func() {
			var details = shortBook.printBookDetails()
			Expect(details).Should(Equal("Book Title: Book Title 1"))
		})
		It("Has an Author", func() {
			Expect(shortBook).Should(HaveAnAuthor("John Doe"))
		})
		It("Has Page count", func() {
			Expect(shortBook).Should(HavePageCount(BeNumerically(">", 0)))
		})
		It("Has Page count", func() {
			Expect(emptyBook).Should(HavePageCount(BeNumerically(">", 0)))
		})

	})

})
