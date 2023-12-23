package router // package router to handle calls at endpoints

import (
	"fmt"      // Go module to format strings
	"net/http" // Go module to handle HTTP requests
)

/**
* Function to greet user at root
 */
func GreetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
