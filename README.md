# Zolo Book Sharing Backend

This repository contains the backend for an application that allows users to share books among themselves, users can add books, borrow a book, return a book and see all available books. 

## Enviornment Variables

- `DB_USER` - To store username of the mysql user (default = "kali")
- `DB_PASS` - To store password of the mysql user (default = "kali")
- `DB_HOST` - To store hostname of the mysql database (default = "127.0.0.1")
- `DB_PORT` - To store port of the mysql database (default = "3306")
- `DB_NAME` - To store the database name (default = "zolo")


## API Setup
1. Install and setup go on your system.
2. Create a MySQL database in your machine and store database details in the enviornment
3. The Database has the following design. 
 ![image](https://github.com/skushagra/zolo-backend/assets/66439372/89d0bafa-7ba2-431b-a705-094609ffb76a)

4. Clone the git repository

```
git clone https://github.com/skushagra/zolo-backend.git
```
5. Get the requirements 
```
go get .
```

## Run
1. Run the server
```
go run .
```

The server will start running on port 9090 by default.

## API Documentation

### 1. GET `/`
Greets user and verifies that the server is running.

### 2. GET `/api/v1/booky`
Returns a list of all the books in the database which are available for sharing.

### 3. PUT `/api/v1/booky`
Adds a new book to the database.
```
Request Payload 
{
	id            int
	title         string
	author        string
	available_till string
	genre         string
	hosted_by      int
}
```

### 4. PUT `/api/v1/booky/<book_id>/borrow`
Borrows a book from the database with the given book id.
```
Request Payload 
{
    id int
    start_time datetime string 
    end_time datetime string
}
```

### 5. POST `/api/v1/booky/<book_id>/borrow/<borrow_id>`
Returns the details of the borrow request with the given borrow id.
