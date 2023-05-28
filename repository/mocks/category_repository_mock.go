package mocks

import (
	"golang-store/model/entity"
	"golang-store/utils"

	mock "github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

// Mock create a new category
func (r *CategoryRepositoryMock) Create(category entity.Category) (*entity.Category, error) {
	args := r.Called(category)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		category := args.Get(0).(entity.Category)
		return &category, nil
	}
}

// Mock update a category
func (r *CategoryRepositoryMock) Update(category entity.Category) (*entity.Category, error) {
	args := r.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

// Mock find category by id
func (r *CategoryRepositoryMock) FindByID(id int) (*entity.Category, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		result := args.Get(0).(entity.Category)
		return &result, nil
	}
}

// Mock find all category
func (r *CategoryRepositoryMock) FindAll(pagination utils.Pagination) (*utils.Pagination, error) {
	args := r.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return &pagination, nil
	}
}
