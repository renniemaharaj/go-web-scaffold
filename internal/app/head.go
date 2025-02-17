package app

import (
	"github.com/renniemaharaj/go-web-scaffold/pkg/elements"
)

func Head() *elements.Head {
	head := elements.Head{}

	head.Title = title

	links := Links()
	head.Links = *links

	scripts := Scripts()
	head.Scripts = *scripts

	metas := make([]elements.Meta, 0)

	metas = append(metas, *elements.MakeMeta("charset", []string{"UTF-8"}, ""))
	metas = append(metas, *elements.MakeMeta("name", []string{"viewport"}, "width=device-width, initial-scale=1.0"))
	metas = append(metas, *elements.MakeMeta("name", []string{"description"}, description))
	metas = append(metas, *elements.MakeMeta("name", []string{"author"}, author))
	metas = append(metas, *elements.MakeMeta("name", []string{"keywords"}, keywords))
	metas = append(metas, *elements.MakeMeta("charset", []string{"UTF-8"}, ""))

	head.Metas = metas

	return &head
}
