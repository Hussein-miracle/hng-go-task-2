package routes

import (
	"github.com/Hussein-miracle/hng-go-task-2/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerEngine *gin.Engine, controller *controllers.PersonController) {

	routerEngine.POST("/api", controller.CreateUser)
	routerEngine.GET("/api/:user_id", controller.GetUser)
	routerEngine.PUT("/api/:user_id", controller.UpdateUser)
	routerEngine.DELETE("/api/:user_id", controller.DeleteUser)
}
