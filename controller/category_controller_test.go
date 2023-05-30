package controller

import (
	"golang-store/model/entity"
	"golang-store/model/web"
	"golang-store/service/mocks"

	mock "github.com/stretchr/testify/mock"
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

var categoryService = &mocks.CategoryServiceMock{Mock: mock.Mock{}}
var categoryController = CategoryControllerImpl{CategoryService: categoryService}

/*func TestCreateSuccess(t testing.T) {
	categoryService.Mock.On("Create", categoryCreateInput).Return(category, nil)
	gin := gin.Default()
	req := httptest.NewRecorder()
	gin.POST("/category", categoryController.Create)

}*/
