package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	//GetAll(c *gin.Context)
}
