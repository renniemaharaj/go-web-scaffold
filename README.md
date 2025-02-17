# Go Web Scaffold

A minimalist web application scaffold using Go and TailwindCSS.

## Getting Started

1. Install dependencies:
```bash
npm i
```

2. Run the development watcher:
```bash
go run pkg/watcher/main.go
```

This will start a file watcher that enables live reload when editing files.

## Project Structure

- `cmd/example/main.go`: Your main application entry point
- `pkg/watcher/main.go`: Development file watcher
- `static/`: Static assets and compiled CSS
- `templates/`: HTML templates

## Development

Edit your application code in `cmd/example/main.go` or create new examples in the `cmd` folder.

Basic example:
```go
package main

import (
	"fmt"

	"github.com/renniemaharaj/go-web-scaffold/internal/app"
)

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
}

```

The watcher will automatically detect changes and rebuild your application. You can build on what i've built, but I plan on integrating typescript, docker, golang linting, writing tests and formatting soon.