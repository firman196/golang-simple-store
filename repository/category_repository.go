package repository

import (
	"golang-store/model/entity"
	"golang-store/utils"
)

type CategoryRepository interface {
	Create(category entity.Category) (*entity.Category, error)
	Update(category entity.Category) (*entity.Category, error)
	FindByID(id int) (*entity.Category, error)
	FindAll(pagination utils.Pagination) (*utils.Pagination, error)
}
