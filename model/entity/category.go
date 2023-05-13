package entity

import "time"

type Category struct {
	Id           int
	Icon         string
	CategoryName string
	Slug         string
	Notes        string
	IsAktif      string
	ParentId     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
