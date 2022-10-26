package internal

import (
	"github.com/faveroferreira/maromba-center/internal/controllers"
	"github.com/gin-gonic/gin"
)

func StartGin(profile string) {
	SetGinMode(profile)

	server := gin.New()
	server.Use(gin.Recovery(), gin.LoggerWithWriter(gin.DefaultWriter))
	server = SetupApplication(server)

	if err := server.Run(); err != nil {
		panic(err)
	}
}

func SetupApplication(server *gin.Engine) *gin.Engine {
	healthController := controllers.HealthController{}

	server.GET("/health", healthController.HealthCheck)

	return server
}

func SetGinMode(profile string) {
	if profile == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
