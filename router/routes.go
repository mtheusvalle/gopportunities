package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mtheusvalle/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		// v1.GET("/opening/:id", handler.GetOpeningByIDHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		// v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		// v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}
}
