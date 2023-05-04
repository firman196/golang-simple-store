package service

import (
	"golang-graphql/utils"
	"golang-store/models/entity"
	"golang-store/models/web"
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

func (service *CategoryServiceImpl) Create(input web.CategoryCreateInput) (entity.CategoryRepository, error) {
	category := entity.UserRepository{
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

func (service *CategoryServiceImpl) Update(input web.CategoryUpdateInput) (entity.CategoryRepository, error) {
	category, err := service.repository.FindByID(input.Id)
	if err != nil {
		return nil, err
	}

	if input.Icon != nil {
		category.Icon = input.Icon
	}

	newCategory := entity.UserRepository{
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

func (service *CategoryServiceImpl) GetById(id int) (*entity.CategoryRepository, error) {
	category, err := service.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (service *CategoryServiceImpl) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	categories, err := service.repository.FindAll(pagination)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
