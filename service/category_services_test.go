package service

import (
	"errors"
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/repository/mocks"
	"golang-store/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	Id:           1,
	Icon:         "icon.PNG",
	CategoryName: "Food",
	Slug:         "food",
	Notes:        "notes test",
	IsAktif:      "1",
}

var pagination = utils.Pagination{
	Limit:      10,
	Page:       1,
	Sort:       "desc",
	TotalRows:  100,
	TotalPages: 2,
	Rows: []map[string]interface{}{
		{"id": 1, "icon": "icon.PNG", "category_name": "Food", "slug": "food", "notes": "notes test", "is_aktif": "1"},
		{"id": 2, "icon": "icon.PNG", "category_name": "Food", "slug": "food", "notes": "notes test", "is_aktif": "1"},
		{"id": 3, "icon": "icon.PNG", "category_name": "Food", "slug": "food", "notes": "notes test", "is_aktif": "1"},
	},
}

// Scenario successfully
// testing Create category service using testify and mock
func TestCreateSuccess(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	//positive case
	categoryRepository.Mock.On("Create", category).Return(category, nil)
	category, err := categoryService.Create(categoryCreateInput)

	assert.Nil(t, err)
	assert.NotNil(t, category)
	categoryRepository.Mock.AssertExpectations(t)
}

// Scenario failed
// testing Create category service using testify and mock
func TestFailedCreate(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("Create", category).Return(nil, errors.New("Internal Server Error"))
	categoryErr, err := categoryService.Create(categoryCreateInput)

	assert.Error(t, err)
	assert.Nil(t, categoryErr)
	categoryRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing GetById category using service using testify and mock
func TestFindByIdSuccess(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("FindByID", 1).Return(category, nil)
	category, err := categoryService.GetById(1)

	assert.Nil(t, err)
	assert.NotNil(t, category)
	categoryRepository.Mock.AssertExpectations(t)
}

// Scenario failed category not found
// testing GetById category using service using testify and mock
func TestFindByIdErrNotFound(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("FindByID", 2).Return(nil, errors.New("category with id 1 is not found")).Once()
	errCategory, err := categoryService.GetById(2)
	assert.Error(t, err)
	assert.Nil(t, errCategory)
	categoryRepository.Mock.AssertExpectations(t)
}

// Scenario successfully
// testing Update category using service using testify and mock
func TestUpdateSuccess(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("Update", category).Return(category, nil)
	categoryRepository.Mock.On("FindByID", 1).Return(category, nil)
	category, err := categoryService.Update(categoryUpdateInput)

	assert.Nil(t, err)
	assert.NotNil(t, category)
}

// Scenario failed because user not found
// testing Update category using service using testify and mock
func TestUpdateFailed(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("Update", category).Return(nil, "category with id 1 is not found")
	categoryRepository.Mock.On("FindByID", 1).Return(nil, errors.New("category with id 2 is not found"))
	categoryUpdateErr, err := categoryService.Update(categoryUpdateInput)

	assert.Error(t, err)
	assert.Nil(t, categoryUpdateErr)
}

// Scenario successfully
// testing get all category using service using testify and mock
func TestGetAllSuccess(t *testing.T) {
	var categoryRepository = &mocks.CategoryRepositoryMock{Mock: mock.Mock{}}
	var categoryService = CategoryServiceImpl{repository: categoryRepository}
	categoryRepository.Mock.On("FindAll", pagination).Return(pagination, nil)
	categoryAll, err := categoryService.GetAll(pagination)

	assert.Nil(t, err)
	assert.NotNil(t, categoryAll)
}
