package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-Burak-Atak/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-Burak-Atak/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-Burak-Atak/infrastructure"
)

var bookRepo *book.BookRepository

// Creates db and bookRepo, If table does not exist reads csv file and creates db tables
func init() {
	db := infrastructure.NewPostgresDB("host=localhost user=postgres password=postgres dbname=library port=5432 sslmode=disable")

	bookRepo = book.NewRepository(db)
	bookRepo.Migration()
	allBooks := bookRepo.FindAll()

	if len(allBooks) == 0 {
		csvSlice := helper.ReadCsv("books.csv")
		for _, c := range csvSlice[1:] {
			pagenumber, _ := strconv.Atoi(c[3])
			price, _ := strconv.Atoi(c[4])
			stockquantity, _ := strconv.Atoi(c[5])

			newBook := book.NewModel(c[0], c[1], c[2], pagenumber, price, stockquantity, c[6])
			bookRepo.Create(*newBook)
		}
	}
}

func main() {

	input := os.Args

	// Lists all of books.
	if input[1] == "list" {
		if len(input) == 2 {
			allBooks := bookRepo.FindAll()
			helper.PrintResults(allBooks)
		} else {
			fmt.Println("Wrong input.")
		}

		// Searches in book name, stock code, author name and isbn and prints them.
	} else if input[1] == "search" {
		searchWord := strings.Join(input[2:], " ")
		allBooks := bookRepo.SearchByInput(searchWord)

		if allBooks[0].BookName == "" {
			fmt.Println("The book you are looking for does not exist")
		} else {
			helper.PrintResults(allBooks)
		}

		// Decreases book stock quantity if all conditions are True
	} else if input[1] == "buy" {
		_, err := strconv.Atoi(input[2])
		number, err2 := strconv.Atoi(input[3])
		if err != nil || err2 != nil {
			fmt.Println("Wrong input.")
		} else {
			book, err := bookRepo.SearchById(input[2])

			if err != nil {
				fmt.Println(err)
			} else {
				if number > 0 && number <= book.StockQuantity {
					bookRepo.Buy(book, number)
				} else {
					fmt.Printf("Enter a number beetween 0-%d", book.StockQuantity)
				}
			}
		}

		// Deletes book if book exist
	} else if input[1] == "delete" && len(input) == 3 {
		_, err := strconv.Atoi(input[2])
		if err != nil {
			fmt.Println("Wrong input.")
		} else {
			book, err := bookRepo.SearchById(input[2])
			if err != nil {
				fmt.Println(err)
			} else {
				bookRepo.Delete(book)
			}
		}

	} else {
		fmt.Print("Enter 'list' for listing.\nEnter 'search -words-' for search in books.\nEnter 'buy -id- -number of books-' to buy book.\nEnter 'delete -id-' to delete a book.")
	}
}
