package routes

import (
	"github.com/Hussein-miracle/hng-go-task-2/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerEngine *gin.Engine, controller controllers.UserController) {

	routerEngine.POST("/api", controller.CreateUser)
}
