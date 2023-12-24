package db // package db to handle database calls

import (
	"database/sql" // Go module to handle SQL calls
	"os"           // Go module to handle OS calls

	_ "github.com/go-sql-driver/mysql" // Go module to handle MySQL calls
)

/**
* Book struct to store book details (used for displaying books)
 */
type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	AvailableTill string `json:"available_till"`
	Genre         string `json:"genre"`
	HostedBy      string `json:"hosted_by"`
	Available     int    `json:"available"`
}

/**
* PutBook struct to store book details (used for adding books)
 */
type PutBook struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	AvailableTill string `json:"available_till"`
	Genre         string `json:"genre"`
	HostedBy      int    `json:"hosted_by"`
}

/**
* Borrower struct to store borrower details
 */
type Borrower struct {
	Id        int    `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

/**
* Function to get all books available for sharing
 */
func AllBooks() []Book {
	db, err := Connect()
	if err != nil {
		panic("AllBooks() :-> Error connecting to database")
	}
	defer db.Close()

	results, err := db.Query("SELECT b.ID, b.BOOK_NAME, b.BOOK_AUTHOR, b.AVAILABLE_TILL, b.GENRE, l.FULL_NAME, b.AVAILABLE FROM books b JOIN lenders l ON b.HOSTED_BY = l.ID WHERE b.AVAILABLE = 1;")
	if err != nil {
		panic("AllBooks() :-> Error querying database")
	}

	var books []Book

	for results.Next() {
		var book Book
		err = results.Scan(&book.Id, &book.Title, &book.Author, &book.AvailableTill, &book.Genre, &book.HostedBy, &book.Available)
		if err != nil {
			panic("AllBooks() :-> Error scanning results")
		}
		books = append(books, book)
	}
	return books
}

/**
* Function to add a book for sharing by a user
* @param book Book to be added
 */
func AddBook(book PutBook) {

	db, err := Connect()
	if err != nil {
		panic("AddBook() :-> Error connecting to database")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO books (BOOK_NAME,BOOK_AUTHOR,AVAILABLE_TILL,GENRE,HOSTED_BY) VALUES (?,?,?,?,?)", book.Title, book.Author, book.AvailableTill, book.Genre, book.HostedBy)
	if err != nil {
		panic("AddBook() :-> Error inserting into database")
	}
}

/**
* Function to get all books shared by a user
* @param id ID of the user
* @return Book Books with the given ID
 */
func GetBook(id int) (Book, error) {

	db, err := Connect()
	if err != nil {
		panic("GetBook() :-> Error connecting to database")
	}
	defer db.Close()

	var book Book
	err = db.QueryRow("SELECT b.ID, b.BOOK_NAME, b.BOOK_AUTHOR, b.AVAILABLE_TILL, b.GENRE, l.FULL_NAME, b.AVAILABLE FROM books b JOIN lenders l ON b.HOSTED_BY = l.ID WHERE b.ID = ?;", id).Scan(&book.Id, &book.Title, &book.Author, &book.AvailableTill, &book.Genre, &book.HostedBy, &book.Available)
	if err != nil {
		panic("GetBook() :-> No such book")
	}
	return book, nil
}

/**
* Function to get all books shared by a user
* @param id ID of the book to be borrowed
* @param borrower Borrower details
 */
func BorrowBook(id int, borrower Borrower) {
	db, err := Connect()
	if err != nil {
		panic("BorrowBook() -> Error connecting to database")
	}
	defer db.Close()

	_, err = db.Exec("UPDATE books SET AVAILABLE = 0 WHERE ID = ?", id)
	if err != nil {
		panic("BorrowBook() -> Error while setting book to unavailable")
	}

	_, err = db.Exec("INSERT INTO borrows (BOOK, TAKEN_BY, BORROW_START_TIME, BORROW_END_TIME, BORROW_COMPLETE) VALUES (?, ?, ?, ?, 0)", id, borrower.Id, borrower.StartTime, borrower.EndTime)
	if err != nil {
		panic("BorrowBook() -> Error while inserting into borrows table")
	}
}

/**
* Function to get all books shared by a user
* @param id ID of the borrow
 */
func ReturnBook(borrow_id int) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("UPDATE books SET AVAILABLE = 1 WHERE ID = (SELECT BOOK FROM borrows WHERE ID = ?)", borrow_id)
	if err != nil {
		panic("ReturnBook() -> Error while setting book to available")
	}

	_, err = db.Exec("UPDATE borrows SET BORROW_COMPLETE = 1 WHERE ID = ?", borrow_id)
	if err != nil {
		panic("ReturnBook() -> Error while setting borrow to complete")
	}
}

/**
* Function to connect to the database
* @return db Database connection
* @return error Error while connecting to database
 */
func Connect() (*sql.DB, error) {
	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		db_user = "kali"
	}

	db_pass := os.Getenv("DB_PASS")
	if db_pass == "" {
		db_pass = "kali"
	}

	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		db_host = "127.0.0.1"
	}

	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		db_port = "3306"
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		db_name = "zolo"
	}

	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_name)
	if err != nil {
		return nil, err
	}
	return db, nil
}
