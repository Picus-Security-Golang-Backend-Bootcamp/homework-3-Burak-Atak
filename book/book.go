package book

import (
	"errors"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	BookName      string `json:"bookName"`
	StockCode     string `json:"stockCode"`
	Isbn          string `json:"isbn"`
	PageNumber    int    `json:"pageNumber"`
	Price         int    `json:"price"`
	StockQuantity int    `json:"stockQuantity"`
	Author        string `json:"author"`
}

type BookRepository struct {
	db *gorm.DB
}

// Creates new book model
func NewModel(bookname string, stockcode string, isbn string, pagenumber int, price int, stockquantity int, author string) *Books {
	return &Books{
		BookName:      bookname,
		StockCode:     stockcode,
		Isbn:          isbn,
		PageNumber:    pagenumber,
		Price:         price,
		StockQuantity: stockquantity,
		Author:        author,
	}
}

// Creates book repositors
func NewRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

// Migration
func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Books{})
}

// Creates new book in database
func (r *BookRepository) Create(book Books) {
	r.db.Create(&book)
}

// Deletes book in db
func (r *BookRepository) Delete(b Books) {
	r.db.Delete(&b)
}

// Decrease stock quantity and saves new book
func (r *BookRepository) Buy(b Books, c int) {
	b.StockQuantity -= c
	r.db.Save(b)

}

// Finds all books in db
func (r *BookRepository) FindAll() []Books {
	var books []Books
	r.db.Find(&books)

	return books
}

// Search the book by id and returns book if it exist
func (r *BookRepository) SearchById(id string) (Books, error) {
	var book Books
	r.db.Find(&book, "id="+id)

	if book.BookName == "" {
		return book, errors.New("Book couldn't found")
	}
	return book, nil
}

// Searches in stock code, book name, isbn and author, returns all matched books
func (r *BookRepository) SearchByInput(name string) []Books {
	var books []Books
	r.db.Find(&books,
		"LOWER(book_name) LIKE LOWER('%"+name+"%')"+
			"OR LOWER(stock_code) LIKE LOWER('%"+name+"%')"+
			"OR LOWER(isbn) LIKE LOWER('%"+name+"%')"+
			"OR LOWER(author) LIKE LOWER('%"+name+"%')")

	return books
}

// Updates book in db
func (r *BookRepository) Update(b Books) {
	r.db.Save(&b)
}
