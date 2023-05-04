package service

import (
	"golang-graphql/utils"
	"golang-store/model/entity"
	"golang-store/models/web"
)

type CategoryService interface {
	Create(input web.CategoryCreateInput) (entity.Category, error)
	Update(input web.CategoryUpdateInput) (entity.Category, error)
	GetById(id int) (entity.Category, error)
	GetAll(pagination utils.Pagination) (*utils.Pagination, error)
}
