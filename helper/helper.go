package helper

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-Burak-Atak/book"
)

// Prints given book slice
func PrintResults(allBooks []book.Books) {
	for i := 0; i < len(allBooks); i++ {
		fmt.Printf("ID: %d Book: %s Author: %s Price: %d\n", allBooks[i].ID, allBooks[i].BookName, allBooks[i].Author, allBooks[i].Price)
	}
}
