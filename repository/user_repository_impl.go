package repository

import (
	"golang-store/model/entity"
	"golang-store/utils"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user entity.User) (*entity.User, error) {
	err := r.db.Model(&user).Where("email = ?", user.Email).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var users []entity.User
	if err := r.db.Scopes(pagination.Paginate(users, r.db)).Find(&users).Error; err != nil {
		return nil, err
	}
	pagination.Rows = users
	return &pagination, nil
}
