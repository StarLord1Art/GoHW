package library

type Library struct {
	books      map[string]int
	storage    Storage
	generateID func(...int) int
}

func (lib *Library) Initialize(generateID func(...int) int, isMapStorageChosen bool) {
	lib.books = make(map[string]int)
	if isMapStorageChosen {
		lib.storage = Storage{booksMp: make(map[int]Book)}
	} else {
		lib.storage = Storage{booksSl: make([]Book, 0)}
	}
	lib.generateID = generateID
}

func (lib *Library) ChangeIDGenerator(generateID func(...int) int) {
	lib.generateID = generateID
}

func (lib *Library) ChangeStorage() {
	if lib.storage.isMapStorageChosen {
		lib.books = make(map[string]int)
		lib.storage.booksMp = make(map[int]Book)
		lib.storage.isMapStorageChosen = false
	} else {
		lib.books = make(map[string]int)
		lib.storage.booksSl = make([]Book, 0)
		lib.storage.isMapStorageChosen = true
	}
}

func (lib *Library) GetBook(title string) Book {
	return lib.storage.getBook(lib.books[title])
}

func (lib *Library) AddBook(book *Book) {
	book.setID(lib.generateID(len(lib.books)))
	lib.books[book.Title] = book.getID()
	lib.storage.addBook(book)
}
