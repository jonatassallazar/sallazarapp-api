package routers

import (
	"api/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(g *gin.RouterGroup, gc *controllers.GeneralController) {
	usersRoutes := g.Group("/users")
	{
		// usersRoutes.GET("/login", gc.UserLogin)
		usersRoutes.POST("/signup", gc.UserSignup)
		usersRoutes.GET("/:id", gc.GetUser)
	}
}
