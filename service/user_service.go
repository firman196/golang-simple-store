package service

import (
	"golang-store/model/entity"
	"golang-store/model/web"
)

type UserService interface {
	Register(input web.CreateUser) (*entity.User, error)
	Login(input web.Login) (*web.Token, error)
	RefreshToken(refreshToken string) (*web.Token, error)
}
