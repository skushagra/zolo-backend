package db

import (
	"database/sql" // Go module to handle SQL calls
	// Go module to handle errors
	"os" // Go module to handle OS calls

	_ "github.com/go-sql-driver/mysql" // Go module to handle MySQL calls
)

/**
* Function to setup database
 */

func Setup() {

	db, err := connect()
	if err != nil {
		panic("Error connecting to database")
	}
	defer db.Close()

	create_lenders_table := "create table if not exists lenders (ID int not null auto_increment,FULL_NAME varchar(50) not null,primary key (ID));"
	_, err = db.Exec(create_lenders_table)
	if err != nil {
		panic("Error creating lenders table -> " + err.Error())
	}

	create_lenders := "insert into lenders (FULL_NAME) values ('Kushagra S');"
	_, err = db.Exec(create_lenders)
	if err != nil {
		panic("Error creating lenders -> " + err.Error())
	}

	create_lenders = "insert into lenders (FULL_NAME) values ('Brijanya S');"
	_, err = db.Exec(create_lenders)
	if err != nil {
		panic("Error creating lenders -> " + err.Error())
	}

	create_books_table := "CREATE TABLE IF NOT EXISTS books (ID INT NOT NULL AUTO_INCREMENT,BOOK_NAME VARCHAR(100) NOT NULL,BOOK_AUTHOR varchar(100) not null,AVAILABLE_TILL datetime not null,GENRE TEXT not null,HOSTED_BY INT,PRIMARY KEY (ID),FOREIGN KEY (HOSTED_BY) REFERENCES lenders(ID));"
	_, err = db.Exec(create_books_table)
	if err != nil {
		panic("Error creating books table")
	}

	create_borrowers_table := "CREATE TABLE IF NOT EXISTS borrows (ID int not null auto_increment,BOOK int,TAKEN_BY int,BORROW_START_TIME datetime not null,BORROW_END_TIME datetime not null,BORROW_COMPLETE int not null default 0,primary key (ID),foreign key (BOOK) references books(ID),foreign key (TAKEN_BY) references lenders(ID));"
	_, err = db.Exec(create_borrowers_table)
	if err != nil {
		panic("Error creating borrowers table")
	}

}

func connect() (*sql.DB, error) {
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
