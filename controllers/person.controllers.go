package controllers

import (
	"net/http"
	"time"

	"github.com/Hussein-miracle/hng-go-task-2/models"
	"github.com/Hussein-miracle/hng-go-task-2/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonController struct {
	PersonService services.PersonService
}

func NewUserController(uc services.PersonService) PersonController {
	return PersonController{
		PersonService: uc,
	}
}

func (PersonControllerPointer *PersonController) CreateUser(ctx *gin.Context) {
	var user models.Person

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err := PersonControllerPointer.PersonService.CreateUser(&user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "person created successfully", "status": "success", "data": user})

}

func (PersonControllerPointer *PersonController) GetUser(ctx *gin.Context) {
	param := ctx.Param("user_id")

	user, err := PersonControllerPointer.PersonService.GetUser(param)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}

func (PersonControllerPointer *PersonController) DeleteUser(ctx *gin.Context) {
	param := ctx.Param("user_id")

	err := PersonControllerPointer.PersonService.DeleteUser(param)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadGateway, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Person Deleted Successfully", "status": "success"})
}

func (PersonControllerPointer *PersonController) UpdateUser(ctx *gin.Context) {
	var user *models.Person

	param := ctx.Param("user_id")

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := PersonControllerPointer.PersonService.UpdateUser(user, param)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "success", "message": "Person updated successfully"})
}

func (PersonControllerPointer *PersonController) CronJ(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "server woke up", "status": "success"})
}
