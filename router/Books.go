package router // package router to handle calls at endpoints

import (
	"encoding/json"   // Go module to encode and decode JSON
	"net/http"        // Go module to handle HTTP requests
	"strconv"         // Go module to convert strings to int and vice versa
	"strings"         // Go module to handle strings to process different endpoints
	"zolo/backend/db" // Go module to handle database calls
)

/**
* Function to show all books available for sharing
* @api {get} /api/v1/booky/ Get all books
 */
func showAllBooks(w http.ResponseWriter, r *http.Request) {
	var books = db.AllBooks() // get all books from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

/**
* Function to add a book for sharing by a user
* @api {put} /api/v1/booky/ Add a book
* @apiParam {String} title Title of the book
* @apiParam {String} author Author of the book
* @apiParam {String} available_till Date till which the book is available for sharing
* @apiParam {String} genre Genre of the book
* @apiParam {String} hosted_by ID of the user who is sharing the book
 */
func addBook(w http.ResponseWriter, r *http.Request) {
	var book db.PutBook
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		panic("addBook() :-> Error decoding book")
	}
	db.AddBook(book) // add book to database
}

/**
* Function to borrow a book by a user
* @api {put} /api/v1/booky/<book_id>/borrow Borrow a book
* @apiParam {String} id ID of the user who is borrowing the book
* @apiParam {String} start_time Date and time when the book is borrowed
* @apiParam {String} end_time Date and time when the book is returned
 */
func borrowBook(w http.ResponseWriter, r *http.Request, book_to_borrow string) {
	book_id, err := strconv.Atoi(book_to_borrow)
	var borrower db.Borrower
	err = json.NewDecoder(r.Body).Decode(&borrower)
	if err != nil {
		panic("borrowBook() :-> Error decoding borrower")
	}
	book, err := db.GetBook(book_id)
	if err != nil {
		panic("borrowBook() :-> Error getting book")
	}
	if book.Available == 1 {
		db.BorrowBook(book_id, borrower) // borrow book from database if available
	} else {
		panic("borrowBook() :-> Book not available")
	}

}

/**
* Function to return a borrowed book by a user
* @api {post} /api/v1/booky/<book_id>/borrow/<borrow_id> Return a book
 */
func returnBook(w http.ResponseWriter, r *http.Request, borrow_id string) {
	borrow_id_int, err := strconv.Atoi(borrow_id)
	if err != nil {
		panic("returnBook() :-> Error converting borrow_id to int")
	}
	db.ReturnBook(borrow_id_int)
}

/**
* Function to process the book related calls
* @api {get} /api/v1/booky/ Get all books
* @api {put} /api/v1/booky/ Add a book
* @api {put} /api/v1/booky/<book_id>/borrow Borrow a book
* @api {post} /api/v1/booky/<book_id>/borrow/<borrow_id> Return a book
 */
func Books(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		showAllBooks(w, r)
	case "PUT":
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) == 3 {
			addBook(w, r)
		}
		if len(pathParts) == 6 && pathParts[5] == "borrow" {
			book_to_borrow := pathParts[4]
			borrowBook(w, r, book_to_borrow)
		}
	case "POST":
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) == 7 && pathParts[5] == "borrow" {
			borrow_id := pathParts[6]
			returnBook(w, r, borrow_id)
		}
	}
}
