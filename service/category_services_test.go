package service

import (
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryServiceImpl{repository: categoryRepository}
var category = entity.Category{
	Icon:         "icon.PNG",
	CategoryName: "Food",
	Slug:         "food",
	Notes:        "notes test",
	IsAktif:      "1",
}

var categoryInput = web.CategoryCreateInput{
	Icon:         "icon.PNG",
	CategoryName: "Food",
	Slug:         "food",
	Notes:        "notes test",
	IsAktif:      "1",
}

func TestPositiveCreateCategory(t *testing.T) {
	categoryRepository.Mock.On("Create", &category).Return(category)
	category, err := categoryService.Create(categoryInput)

	assert.Nil(t, err)
	assert.NotNil(t, category)
}

func TestFindById(t *testing.T) {
	categoryRepository.Mock.On("FindByID", "1").Return(category, nil)
	category, err := categoryService.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, category)
}
