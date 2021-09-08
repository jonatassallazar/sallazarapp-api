package routers

import (
	"api/cmd/api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, gc *controllers.GeneralController) {

	baseApi := r.Group("/v1")
	{
		baseApi.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "v1/index",
			})
		})
		UsersRoutes(baseApi, gc)
	}
}
