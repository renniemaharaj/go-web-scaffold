package app

import (
	"github.com/renniemaharaj/go-web-scaffold/pkg/elements"
)

func Scripts() *[]elements.Script {
	scripts := make([]elements.Script, 0)

	scripts = append(scripts, elements.Script{
		Src: "/static/script.js",
	})

	return &scripts
}
