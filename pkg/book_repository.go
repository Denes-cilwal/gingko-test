package pkg

type BookRepository interface {
	GetBooks() ([]string, error)
}
