package main

import (
	"GoHW1/library"
	"fmt"
)

func main() {
	books := []library.Book{library.Book{Title: "War and Peace", Author: "Lev Tolstoy"},
		library.Book{Title: "Romeo and Juliet", Author: "William Shakespeare"},
		library.Book{Title: "To the Moon and back", Author: "Jules Vern"},
	}

	lib := library.Library{}
	lib.Initialize()
	for _, book := range books {
		lib.AddBook(&book)
	}

	fmt.Println(lib.GetBook("War and Peace"))
	fmt.Println(lib.GetBook("Romeo and Juliet"))

	lib.Clear()
	lib.AddBook(&library.Book{Title: "Don Quixote", Author: "Miguel de Cervantes"})
	lib.AddBook(&library.Book{Title: "The Three Musketeers", Author: "Alexandre Dumas"})
	lib.AddBook(&library.Book{Title: "Robinson Crusoe", Author: "Daniel Defoe"})

	fmt.Println(lib.GetBook("Robinson Crusoe"))
	fmt.Println(lib.GetBook("The Three Musketeers"))
	fmt.Println(lib.GetBook("War and Peace"))
}
