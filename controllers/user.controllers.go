package controllers

import (
	"net/http"

	"github.com/Hussein-miracle/hng-go-task-2/models"
	"github.com/Hussein-miracle/hng-go-task-2/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(uc services.UserService) UserController {
	return UserController{
		UserService: uc,
	}
}

func (UserControllerPointer *UserController) CreateUser(ctx *gin.Context) {
	var user *models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

}
