package routes

import (
	"note/app/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", handler.RegisterUser)
}