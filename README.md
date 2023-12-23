# Zolo Book Sharing Backend
created by Kushagra

## Setup
1. Install and setup go on your system
2. Clone the git repository 
```
git clone https://github.com/skushagra/zolo-backend.git
```
3. Get the requirements 
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
