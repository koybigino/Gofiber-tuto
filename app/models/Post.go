package models

import "time"

type PostBase struct {
	Title     string `json:"title" validate:"required,min=10,max=20"`
	Content   string `json:"content"  validate:"required,min=20"`
	Published bool   `json:"published"  validate:"required,boolean"`
}

type PostRequest struct {
	PostBase
}

type PostResponse struct {
	Id int `json:"id" gorm:"primaryKey"`
	PostBase
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	Id int `json:"id" gorm:"primaryKey"`
	PostResponse
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ParseToResponse(p Post, pr *PostResponse) {
	pr.Id = p.Id
	pr.Title = p.Title
	pr.Content = p.Content
	pr.Published = p.Published
	pr.CreatedAt = p.CreatedAt
}

func ParseToPost(p *Post, pr PostRequest) {
	p.Title = pr.Title
	p.Content = pr.Content
	p.Published = pr.Published
}
