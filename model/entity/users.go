package entity

import "time"

type User struct {
	Id               int
	Firstname        string
	Lastname         string
	Password         string
	Salt             string
	Address          string
	Image            string
	Email            string
	EmailVerifiedAt  time.Time
	VerificationCode string
	PhoneNumber      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	IsBanned         string
	IsDelete         string
	DeletedAt        time.Time
	DeletedBy        int
}
