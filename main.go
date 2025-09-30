package main

import (
	"GoHW1/library"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func generateID1(...int) int {
	a, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return int(a.Int64())
}

func generateID2(params ...int) int {
	return params[0] + 1
}

func main() {
	books := []library.Book{{Title: "War and Peace", Author: "Lev Tolstoy"},
		{Title: "Romeo and Juliet", Author: "William Shakespeare"},
		{Title: "To the Moon and back", Author: "Jules Vern"},
	}

	lib := library.Library{}
	lib.Initialize(generateID1, true)
	for _, book := range books {
		lib.AddBook(&book)
	}

	fmt.Println(lib.GetBook("War and Peace"))
	fmt.Println(lib.GetBook("Romeo and Juliet"))
	lib.ChangeIDGenerator(generateID2)
	fmt.Println(lib.GetBook("To the Moon and back"))

	lib.ChangeStorage()
	lib.AddBook(&library.Book{Title: "Don Quixote", Author: "Miguel de Cervantes"})
	lib.AddBook(&library.Book{Title: "The Three Musketeers", Author: "Alexandre Dumas"})
	lib.AddBook(&library.Book{Title: "Robinson Crusoe", Author: "Daniel Defoe"})

	fmt.Println(lib.GetBook("Robinson Crusoe"))
	fmt.Println(lib.GetBook("The Three Musketeers"))
	fmt.Println(lib.GetBook("War and Peace"))
}
