package app

import (
	"github.com/renniemaharaj/go-web-scaffold/pkg/elements"
)

func Style() *elements.Style {
	style := elements.Style{}

	style.Selection = "body"
	style.Styles = map[string]string{
		"color": "white",
	}

	return &style
}
