package models

import (
	"os" // Go module to handle OS calls

	_ "github.com/go-sql-driver/mysql" // Go module to handle MySQL calls
	"gorm.io/driver/mysql"             // Go module to handle MySQL calls
	"gorm.io/gorm"                     // Go module to implement ORM
)

/**
* Function to setup database
 */
func Setup() {

	db, err := connect()
	if err != nil {
		panic("Setup() :-> Error connecting to database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&Lenders{},
		&Books{},
		Borrows{},
	)
}

func connect() (*gorm.DB, error) {
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

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
