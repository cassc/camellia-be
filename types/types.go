package types

import "olympos.io/encoding/edn"

type PostMeta struct {
	Title       string
	Layout      edn.Keyword
	Tags        []string
	Toc         bool
	Description string
	Keywords    string
}

type Post struct {
	FileName string
	Meta     PostMeta
	PostDate string
}

type TagCount struct {
	Tag   string
	Count int
}

type Article struct {
	Title   string
	URL     string
	Content string
	Status  int
	Updated int
}
