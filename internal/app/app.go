package app

import "github.com/renniemaharaj/go-web-scaffold/pkg/elements"

var language = "en"
var title = "Document Representer!"
var description = "This is a simple document go representer for building and generating HTML documents."
var author = "Rennie Maharaj"
var keywords = "thewriterco"

func MyDocument() *elements.Document {
	// Create a new doc
	var doc = elements.BlankDocument()

	// Set our document language
	doc.Language = language

	// Set our document head
	doc.Head = *Head()

	// Set our document body
	doc.Body = *Body()

	// Set our document style
	// doc.Head.Styles = append(doc.Head.Styles, *Style())

	// Return our document
	return doc
}
