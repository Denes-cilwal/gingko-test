package main

type Books struct {
	Name    string
	Author  string
	ISBN    string
	Summary string
}

func (b *Books) GetAuthorName() string {
	return b.Author
}

func (b *Books) GetName() string {
	return b.Name
}
