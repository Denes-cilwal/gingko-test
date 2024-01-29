package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
Normally, this code would cause a compilation error in Go because of the direct invocation of Describe.
The correct approach, as discussed earlier, is to use var _ = Describe(...) to comply with Go's rules.
The underscore (_) serves as the blank identifier in Go, indicating that the returned value of Describe is intentionally being ignored.
*/
var _ = Describe("Books", func() {

	It("returns the name of the book", func() {
		book := Books{
			Name: "A Tale of Two Cities",
		}
		Expect(book.GetName()).To(Equal("A Tale of Two Cities"))
	})
	It("returns the author of the book", func() {
		book := Books{
			Author: "Charles Dickens",
		}
		Expect(book.GetAuthorName()).To(Equal("Charles Dickens"))
	})

})

// using before each and (avoid DRY case)

var _ = Describe("Book", func() {
	var book Books

	/*
		If there are multiple test cases (It blocks) within the Describe block,
		this setup process is repeated for each test case, ensuring they all start with a fresh instance of Book and a valid state check.
	*/
	BeforeEach(func() {
		book = Books{
			Name:   "Les Miserables",
			Author: "Victor Hugo",
		}
	})

	It("should have the correct Name", func() {
		Expect(book.Name).To(Equal("Les Miserables"))
	})

	// More test cases...
})

// container nodes use case - Describe, when, context
//
//	Describe different capabilities of your code and explore the behavior of each capability across different Contexts.
var _ = Describe("Books", func() {
	var book Books

	BeforeEach(func() {
		book = Books{
			Author: "Victor Hugo",
		}
	})

	Describe("Extracting the author's first and last name", func() {
		Context("When the author has both names", func() {
			// add test case
		})

		Context("When the author only has one name", func() {
			BeforeEach(func() {
				book.Author = "Hugo"
			})

			It("interprets the single author name as a last name", func() {
				Expect(book.GetAuthorName()).To(Equal("Hugo"))
			})

		})

	})
})
