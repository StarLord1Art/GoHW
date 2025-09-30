package library

type Storage struct {
	booksMp            map[int]Book
	booksSl            []Book
	isMapStorageChosen bool
}

func (storage *Storage) getBook(id int) Book {
	if storage.isMapStorageChosen {
		return storage.booksMp[id]
	} else {
		for _, book := range storage.booksSl {
			if book.getID() == id {
				return book
			}
		}
	}
	return Book{}
}

func (storage *Storage) addBook(book *Book) {
	if storage.isMapStorageChosen {
		storage.booksMp[book.id] = *book
	} else {
		storage.booksSl = append(storage.booksSl, *book)
	}
}
