/**
* Project: Zolo Backend
* Author: Kushagra S
* Created: Sunday, 24th December 2023
 */

package main // root file of the project with package name main

import (
	"net/http"            // Go module to handle HTTP requests
	"zolo/backend/db"     // Go module to handle utility functions
	"zolo/backend/router" // Go module to handle calls at endpoints
)

/**
* Function to run the backend
 */
func runBackend() {
	http.HandleFunc("/", router.GreetUser)          // Greet user at root
	http.HandleFunc("/api/v1/booky/", router.Books) // handle all book related calls
	http.ListenAndServe(":9090", nil)               // listen on port 9090
}

/**
* Function to run the backend
 */
func main() {
	db.Setup()   // setup the database
	runBackend() // run the backend
}
