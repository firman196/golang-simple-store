package repository

import "golang-store/model/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	Update(user entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}
