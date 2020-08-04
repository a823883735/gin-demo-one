package router

import (
	"gin-demo-one/src/controllers"
	"github.com/gin-gonic/gin"
)

var userRouter *gin.RouterGroup

func init() {
	userRouter = Router.Group("/user")
	{
		//获取人员列表
		userRouter.GET("/find", controllers.GetUsers)
		userRouter.POST("/add", controllers.AddUser)
		userRouter.DELETE("/delete", controllers.DeleteUser)
		userRouter.PUT("/updata", controllers.UpdataUser)
		userRouter.GET("/findUserById", controllers.SelectUser)
		userRouter.POST("/login", controllers.Login)
	}
}
