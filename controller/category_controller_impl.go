package controller

import (
	"golang-store/service"
	"net/http"
	"strconv"

	"golang-store/model/web"
	"golang-store/utils"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryControllerImpl(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (service *CategoryControllerImpl) Create(c *gin.Context) {
	var input web.CategoryCreateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	category, errService := service.CategoryService.Create(input)
	if errService != nil {
		response := utils.ApiResponse("Create category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create category success", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
}

func (service *CategoryControllerImpl) Update(c *gin.Context) {
	var input web.CategoryUpdateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	category, errService := service.CategoryService.Update(input)
	if errService != nil {
		response := utils.ApiResponse("Update category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update category success", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
}

func (service *CategoryControllerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Get data category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, errService := service.CategoryService.GetById(id)
	if errService != nil {
		response := utils.ApiResponse("Get data category failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data category success", http.StatusOK, "success", category)
	c.JSON(http.StatusOK, response)
}

/*
func (service *CategoryControllerImpl) GetAll(c *gin.Context) {
	var param utils.Pagination
	err := c.ShouldBind(&param)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Get All Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	categories, errService := service.CategoryService.GetAll(param)

	if errService != nil {
		response := utils.ApiResponse("Get all data category failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get all data category success", http.StatusOK, "success", categories)
	c.JSON(http.StatusOK, response)
}*/
