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

var categoryCreateInput = web.CategoryCreateInput{
	Icon:         "icon.PNG",
	CategoryName: "Food",
	Slug:         "food",
	Notes:        "notes test",
	IsAktif:      "1",
}

var categoryUpdateInput = web.CategoryUpdateInput{
	Icon:         "icon.PNG",
	CategoryName: "Food",
	Slug:         "food",
	Notes:        "notes test",
	IsAktif:      "1",
}

// testing Create category service using testify and mock
func TestCreateCategory(t *testing.T) {
	//positive case
	t.Run("success", func(t *testing.T) {
		categoryRepository.Mock.On("Create", &category).Return(category, nil)
		category, err := categoryService.Create(categoryCreateInput)

		assert.Nil(t, err)
		assert.NotNil(t, category)
	})

	//negative case
	t.Run("error", func(t *testing.T) {
		categoryRepository.Mock.On("Create", &category).Return(nil, "internal server error")
		category, err := categoryService.Create(categoryCreateInput)

		assert.NotNil(t, err)
		assert.Nil(t, category)
	})

}

// testing GetById category using service using testify and mock
func TestFindById(t *testing.T) {
	//positive case
	t.Run("success", func(t *testing.T) {
		categoryRepository.Mock.On("FindByID", "1").Return(category, nil)
		category, err := categoryService.GetById(1)

		assert.Nil(t, err)
		assert.NotNil(t, category)
	})

	//negative case
	t.Run("error", func(t *testing.T) {
		categoryRepository.Mock.On("FindByID", "1").Return(nil, "category with id 1 is not found")
		category, err := categoryService.GetById(1)

		assert.NotNil(t, err)
		assert.Nil(t, category)
	})
}

// testing Update category using service using testify and mock
func TestUpdate(t *testing.T) {
	//positive case
	t.Run("success", func(t *testing.T) {
		categoryRepository.Mock.On("Update", &category).Return(category, nil)
		category, err := categoryService.Update(categoryUpdateInput)

		assert.Nil(t, err)
		assert.NotNil(t, category)
	})

	//negative case
	t.Run("error", func(t *testing.T) {
		categoryRepository.Mock.On("Update", &category).Return(nil, "category with id 1 is not found")
		category, err := categoryService.Update(categoryUpdateInput)

		assert.NotNil(t, err)
		assert.Nil(t, category)
	})
}
