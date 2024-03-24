package routes

import (
	"go-alimentos/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		alimentos := main.Group("alimentos")
		{
			alimentos.GET("/", controllers.Listar_Alimentos)
			alimentos.GET("/:comida", controllers.Buscar_Alimento)
		}
	}

	return router
}
