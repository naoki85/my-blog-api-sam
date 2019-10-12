package model

type Post struct {
	Id          int
	Category    string
	Title       string
	Content     string
	ImageUrl    string
	PublishedAt string
}

type Posts []Post
