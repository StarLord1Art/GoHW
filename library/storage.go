package library

type Storage struct {
	books map[int]Book
}

func (storage *Storage) getBook(id int) Book {
	return storage.books[id]
}

func (storage *Storage) addBook(book *Book) {
	storage.books[book.id] = *book
}
