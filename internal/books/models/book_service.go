package models

type BookService interface {
	CreateBook(book *Book) error
	GetBook(id int64) (*Book, error)
	GetAllBooks() ([]*Book, error)
	UpdatedBook(int int64, book *Book) error
	DeleteBook(int int64) error
}
