package elements

import (
	"fmt"
	"os"

	"github.com/renniemaharaj/go-web-scaffold/pkg/server"
	"github.com/renniemaharaj/go-web-scaffold/pkg/utils"
)

// Document struct represents an entire HTML document.
type Document struct {
	Language string
	Head     Head
	Body     Body
}

// This function creates and returns a blank document skeleton.
func BlankDocument() *Document {
	return &Document{
		Head: Head{
			Title:   "",
			Metas:   []Meta{},
			Links:   []Link{},
			Scripts: []Script{},
		},
		Body: Body{},
	}
}

// This function builds and serves a document on the specified port and directory.
func (doc *Document) BuildAndServe(port, dist string) (chan []byte, chan []byte) {
	// Transform the document to HTML
	doc.Build(dist)

	// Channels for WebSocket communication
	chanS := make(chan []byte)
	chanR := make(chan []byte)

	// Start the HTTP/WebSocket server
	server.WServer(port, dist, chanS, chanR)

	return chanS, chanR
}

// This function transforms a Document and exports it to the file specified. Export as .html
func (doc *Document) Build(dist string) error {
	// Transform the document via markup builder
	markup := doc.BuildMarkup()

	if err := os.MkdirAll(dist, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Write the file to the dist directory
	filename := fmt.Sprintf("%v/index.html", dist)
	if err := os.WriteFile(filename, []byte(*markup), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	// Run build CSS
	go utils.RunBuildCSS()

	return nil
}

func (doc *Document) BuildMarkup() *string {
	markup := "<!DOCTYPE html>\n<html lang=\"" + doc.Language + "\">\n"

	markup += doc.Head.BuildMarkup()
	markup += doc.Body.BuildMarkup()

	markup += "</html>\n"

	return &markup
}
