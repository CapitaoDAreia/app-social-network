package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (p *Post) ValidatePostFields() error {
	if p.Title == "" {
		return errors.New("must have a title field")
	}

	if p.Content == "" {
		return errors.New("must  have a content fiel")
	}

	return nil
}

func (p *Post) FormatPostFields() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}

func (p *Post) PreparePostData() error {
	if err := p.ValidatePostFields(); err != nil {
		return err
	}

	p.FormatPostFields()

	return nil
}
