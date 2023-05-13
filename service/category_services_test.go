package service

import (
	"golang-store/model/web"
	"golang-store/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryServiceImpl{repository: categoryRepository}

func TestPositiveCreateCategory(t *testing.T) {
	category := web.CategoryCreateInput{
		Icon:         "icon.PNG",
		CategoryName: "Food",
		Slug:         "food",
		Notes:        "notes test",
		IsAktif:      "1",
	}

	categoryRepository.Mock.On("Create", &category).Return(category)

	category, err := categoryService.Create(&category)

	assert.Nil(t, err)
	assert.NotNil(t, category)
}
