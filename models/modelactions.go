package models // package db to handle database calls

import (
	_ "github.com/go-sql-driver/mysql" // Go module to handle MySQL calls
)

/**
* Borrower struct to store borrower details
 */
type Borrower struct {
	TAKEN_BY  int    `json:"borrower_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

/**
* Function to get all books available for sharing
 */
func AllBooks() []Books {
	db, err := connect()
	if err != nil {
		return nil
	}

	var books []Books
	result := db.Find(&books)
	if result.Error != nil {
		return nil
	}
	return books
}

/**
* Function to add a book for sharing by a user
* @param book Book to be added
 */
func AddBook(book Books) string {

	db, err := connect()
	if err != nil {
		return "Error connecting to database"
	}

	result := db.Create(&book)
	if result.Error != nil {
		return "Error adding book"
	}

	return "Book added successfully"

}

/**
* Function to get all books shared by a user
* @param id ID of the user
* @return Book Books with the given ID
 */
func GetBook(id int) Books {
	db, err := connect()
	if err != nil {
		return Books{}
	}

	var book Books
	result := db.First(&book, id)
	if result.Error != nil {
		return Books{}
	}

	return book
}

/**
* Function to get all books shared by a user
* @param id ID of the book to be borrowed
* @param borrower Borrower details
 */
func BorrowBook(id int, borrowerId int, startTime string, endTime string) {
	db, err := connect()
	if err != nil {
		panic("Error connecting to database")
	}

	// Update book availability
	err = db.Model(&Books{}).Where("id = ?", id).Update("AVAILABLE", 0).Error
	if err != nil {
		panic("Error while updating book availability")
	}

	// Insert into borrows table
	borrow := Borrows{
		BOOK:              id,
		TAKEN_BY:          borrowerId,
		BORROW_START_TIME: startTime,
		BORROW_END_TIME:   endTime,
		BORROW_COMPLETE:   0,
	}
	err = db.Create(&borrow).Error
	if err != nil {
		panic("Error while inserting into borrows table")
	}
}

/**
* Function to get all books shared by a user
* @param id ID of the borrow
 */
func ReturnBook(borrow_id int) {
	db, err := connect()
	if err != nil {
		panic(err.Error())
	}

	// Get the book id from the borrow
	var borrow Borrows
	db.First(&borrow, borrow_id)
	book_id := borrow.BOOK

	// Update book availability
	err = db.Model(&Books{}).Where("ID = ?", book_id).Update("AVAILABLE", 1).Error
	if err != nil {
		panic("ReturnBook() -> Error while setting book to available")
	}

	// Update borrow to complete
	err = db.Model(&Borrows{}).Where("ID = ?", borrow_id).Update("BORROW_COMPLETE", 1).Error
	if err != nil {
		panic("ReturnBook() -> Error while setting borrow to complete")
	}
}
