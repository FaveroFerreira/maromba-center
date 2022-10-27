package internal

import (
	"github.com/faveroferreira/maromba-center/internal/controllers"
	"github.com/faveroferreira/maromba-center/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func StartGin(profile string) {
	SetGinMode(profile)

	server := gin.New()
	server.Use(gin.Recovery(), gin.LoggerWithWriter(gin.DefaultWriter))
	db_url := viper.Get("databaseUrl").(string)
	database := db.Init(db_url)
	server = SetupApplication(server, database)

	if err := server.Run(); err != nil {
		panic(err)
	}
}

func SetupApplication(server *gin.Engine, db *gorm.DB) *gin.Engine {
	healthController := controllers.HealthController{}
	crud_handler := &controllers.Handler{
		DB: db,
	}

	server.GET("/health", healthController.HealthCheck)

	trainer_routes := server.Group("/trainers")
	trainer_routes.POST("/", crud_handler.AddTrainer)
	trainer_routes.GET("/", crud_handler.GetTrainers)
	trainer_routes.GET("/:id", crud_handler.GetTrainer)
	trainer_routes.PUT("/:id", crud_handler.UpdateTrainer)
	trainer_routes.DELETE("/:id", crud_handler.DeleteTrainer)

	return server
}

func SetGinMode(profile string) {
	if profile == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
