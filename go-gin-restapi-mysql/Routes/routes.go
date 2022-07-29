package Routes

import (
	"golang-workspace/src/go-gin-restapi-mysql/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controller.GetUsers)
		grp1.POST("user", Controller.CreateUser)
		grp1.GET("user/:id", Controller.GetUserByID)
		grp1.PUT("user/:id", Controller.UpdateUser)
		grp1.DELETE("user/:id", Controller.DeleteUser)
	}

	return r

}
