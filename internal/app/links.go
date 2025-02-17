package app

import (
	"github.com/renniemaharaj/go-web-scaffold/pkg/elements"
)

func Links() *[]elements.Link {
	links := make([]elements.Link, 0)

	links = append(links, elements.Link{
		Rel:  "stylesheet",
		Href: "/static/css/tailwind.min.css",
	})

	links = append(links, elements.Link{
		Rel:  "stylesheet",
		Href: "/static/index.css",
	})

	return &links
}
