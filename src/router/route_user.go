package router

import (
	"gin-demo-one/src/controllers"
	"github.com/gin-gonic/gin"
)

var userRouter *gin.RouterGroup

func init() {
	userRouter = Router.Group("/user")
	{
		userRouter.GET("/", controllers.GetUsers)
		userRouter.POST("/", controllers.AddUser)
		userRouter.DELETE("/", controllers.DeleteUser)
		userRouter.PUT("/", controllers.UpdataUser)
		userRouter.GET("/findUserById", controllers.SelectUser)
		userRouter.POST("/login", controllers.Login)
	}
}
