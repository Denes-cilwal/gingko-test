package pkg

type BookService interface {
	GetBooks() ([]string, error)
}
