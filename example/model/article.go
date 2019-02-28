package model

import "time"

type Article struct {
	ID        int  `sql:",pk,unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title    string `schema:"title"`
	Content  string `schema:"content"`
	Username string  `schema:"username"`
}

func NewArticle(title, content, username string) *Article {
	article := Article{
		Title:    title,
		Content:  content,
		Username: username,
	}

	return &article
}

func (m *Article) BeforeInsert() error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *Article) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}