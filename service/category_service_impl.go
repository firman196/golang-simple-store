package service

import (
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/repository"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
}

func NewCategoryServiceImpl(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		repository: repository,
	}
}

func (service *CategoryServiceImpl) Create(input web.CategoryCreateInput) (*entity.Category, error) {
	category := entity.Category{
		Icon:         input.Icon,
		CategoryName: input.CategoryName,
		Slug:         input.Slug,
		Notes:        input.Notes,
		IsAktif:      input.IsAktif,
		ParentId:     input.ParentId,
	}

	response, err := service.repository.Create(category)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CategoryServiceImpl) Update(input web.CategoryUpdateInput) (*entity.Category, error) {
	category, err := service.repository.FindByID(input.Id)
	if err != nil {
		return nil, err
	}

	if input.Icon != "" {
		category.Icon = input.Icon
	}

	newCategory := entity.Category{
		Id:           category.Id,
		Icon:         category.Icon,
		CategoryName: input.CategoryName,
		Slug:         input.Slug,
		Notes:        input.Notes,
		IsAktif:      input.IsAktif,
		ParentId:     input.ParentId,
	}

	response, err := service.repository.Update(newCategory)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CategoryServiceImpl) GetById(id int) (*entity.Category, error) {
	category, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

/*
func (service *CategoryServiceImpl) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	categories, err := service.repository.FindAll(pagination)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
*/
