package mocks

import (
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/utils"

	"github.com/stretchr/testify/mock"
)

type CategoryServiceMock struct {
	mock.Mock
}

// Mock create a new category
func (r *CategoryServiceMock) Create(category web.CategoryCreateInput) (*entity.Category, error) {
	args := r.Called(category)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		category := args.Get(0).(entity.Category)
		return &category, nil
	}
}

// Mock update a category
func (r *CategoryServiceMock) Update(category web.CategoryUpdateInput) (*entity.Category, error) {
	args := r.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

// Mock get category by id
func (r *CategoryServiceMock) GetById(id int) (*entity.Category, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

// Mock get all category
func (r *CategoryServiceMock) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	args := r.Called(pagination)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return &pagination, nil
	}
}
