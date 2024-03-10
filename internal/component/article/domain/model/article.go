package model

type Article struct {
	ID      string
	Title   string
	Content string
	UserID  string

	User *ArticleUser
}

type ArticleUser struct {
	ID   string
	Name string
}
