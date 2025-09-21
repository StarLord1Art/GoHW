package library

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Library struct {
	books   map[string]int
	storage Storage
}

func (lib *Library) generateID() int {
	a, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return int(a.Int64())
}

func (lib *Library) Initialize() {
	lib.books = make(map[string]int)
	lib.storage = Storage{books: make(map[int]Book)}
}

func (lib *Library) Clear() {
	lib.books = make(map[string]int)
	lib.storage.books = make(map[int]Book)
}

func (lib *Library) GetBook(title string) Book {
	return lib.storage.getBook(lib.books[title])
}

func (lib *Library) AddBook(book *Book) {
	book.setID(lib.generateID())
	lib.books[book.Title] = book.getID()
	lib.storage.addBook(book)
}
