package service

import (
	"errors"
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/repository"
	"golang-store/utils"
	"os"
	"strconv"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (s *UserServiceImpl) Register(input web.CreateUser) (*entity.User, error) {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		Firstname:   input.Firstname,
		Lastname:    input.Lastname,
		Password:    hash,
		Address:     input.Address,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
	}
	response, err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func (service *UserServiceImpl) Login(input *web.Login) (*web.Token, error) {
	user, err := service.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	auth := utils.CheckPasswordMatch(user.Password, input.Password)

	if auth == false {
		return nil, errors.New("Email & Password not match")
	}

	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))

}

func (service *UserServiceImpl) RefreshToken(refreshToken string) (*web.Token, error) {

}
