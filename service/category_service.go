package service

import (
	"golang-store/model/entity"
	"golang-store/models/web"
)

type CategoryService interface {
	Create(input web.CategoryCreateInput) (entity.Category, error)
	Update(input web.CategoryUpdateInput) (entity.Category, error)
}
