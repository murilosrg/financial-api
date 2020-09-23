package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/internal/controllers"
)

func SetupApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/accounts", controllers.PostAccount)
	api.GET("/accounts/:accountId", controllers.FindAccount)
	api.POST("/transactions", controllers.PostTransaction)
}
