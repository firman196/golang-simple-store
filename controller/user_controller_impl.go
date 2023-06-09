package controller

import (
	"golang-store/model/web"
	"golang-store/service"
	"golang-store/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// RegisterUser		godoc
// @Summary			Register a user
// @Description		Save user data in Db.
// @Param			user body web.CreateUser true "Create user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} utils.Response
// @Router			/api/v1/user [post]
func (service *UserControllerImpl) Register(c *gin.Context) {
	var input web.CreateUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	user, errService := service.UserService.Register(input)
	if errService != nil {
		response := utils.ApiResponse("Register user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Register user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (service *UserControllerImpl) Login(c *gin.Context) {
	var input web.Login

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, errService := service.UserService.Login(input)
	if errService != nil {
		response := utils.ApiResponse("Login user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Login user success", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

func (service *UserControllerImpl) RefreshToken(c *gin.Context) {

}
