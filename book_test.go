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
