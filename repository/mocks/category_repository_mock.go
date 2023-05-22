package mocks

import (
	"golang-store/model/entity"

	mock "github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// Mock create a new category
func (r *CategoryRepositoryMock) Create(category entity.Category) (*entity.Category, error) {
	args := r.Mock.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		category := args.Get(0).(entity.Category)
		return &category, nil
	}
}

// Mock update a category
func (r *CategoryRepositoryMock) Update(category entity.Category) (*entity.Category, error) {
	args := r.Mock.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

// find category by id
func (r *CategoryRepositoryMock) FindByID(id int) (*entity.Category, error) {
	args := r.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

/*
func (r *CategoryRepositoryMock) FindAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var categories []entity.Category
	if err := r.db.Scopes(pagination.Paginate(categories, r.db)).Find(&categories).Error; err != nil {
		return nil, err
	}

	pagination.Rows = categories
	return &pagination, nil
}*/