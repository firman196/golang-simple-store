package entity

import "time"

type Category struct {
	Id           int
	Icon         string
	CategoryName string
	Slug         string
	Notes        string
	IsAktif      string
	ParentId     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
