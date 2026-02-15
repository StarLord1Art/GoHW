package library

type Book struct {
	Title  string
	Author string
	id     int
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) getID() int {
	return b.id
}

func (b *Book) setID(id int) {
	b.id = id
}
