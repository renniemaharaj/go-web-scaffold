package main

import (
	// "fmt"

	"fmt"

	"github.com/renniemaharaj/go-web-scaffold/internal/app"
)

// The port to run the HTTP server on
// var port = "8080"

// The directory to export the document to
var dist = "static"

func main() {
	// Create a new document
	doc := app.MyDocument()

	// Build the document and capture any error
	if err := doc.Build(dist); err != nil {
		fmt.Printf("Error building document: %v\n", err)
		return
	}

	fmt.Println("Document built successfully")
	// fmt.Scanln() // Wait for user input before closing
}
