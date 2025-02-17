package app

import (
	"github.com/renniemaharaj/go-web-scaffold/pkg/elements"
)

func Body() *elements.Body {
	body := elements.Body{}

	// Title (Animated)
	h1 := elements.Element{
		Tag: "h1",
		Attributes: []elements.Attribute{
			{Name: "class", Value: "mt-10 text-center animate-bounce text-4xl font-bold"},
			{Name: "innerHTML", Value: "Go Represent!!"},
			{Name: "id", Value: "title"},
		},
	}

	// Subtitle
	p := elements.Element{
		Tag: "p",
		Attributes: []elements.Attribute{
			{Name: "class", Value: "text-center text-lg text-white mt-4"},
			{Name: "innerHTML", Value: "A Golang-powered website builder that transforms structured data into responsive web pages."},
		},
	}

	// Button (Triggers JS)
	button := elements.Element{
		Tag: "button",
		Attributes: []elements.Attribute{
			{Name: "class", Value: "relative top-25 mt-6 px-6 py-2 bg-green-500 left-[50%] translate-x-[-50%] text-white font-semibold rounded-lg shadow-md hover:bg-green-600 focus:outline-none"},
			{Name: "innerHTML", Value: "Click to Change Title"},
			{Name: "id", Value: "changeTextBtn"},
		},
	}

	// Append elements to body
	body.Elements = append(body.Elements, h1, p, button)

	return &body
}
