package internal

import (
	"github.com/faveroferreira/maromba-center/internal/controllers"
	"github.com/faveroferreira/maromba-center/internal/db"
	"github.com/faveroferreira/maromba-center/internal/services"
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
	trainersController := &controllers.TrainerController{
		Services: &services.TrainerServices{
			DB: db,
		},
	}

	server.GET("/health", healthController.HealthCheck)

	trainers_routes := server.Group("/trainers")
	trainers_routes.POST("/", trainersController.CreateTrainer)
	trainers_routes.GET("/:id", trainersController.GetTrainerById)
	trainers_routes.GET("/", trainersController.GetAllTrainers)
	trainers_routes.PUT("/:id", trainersController.UpdateTrainer)
	trainers_routes.DELETE("/:id", trainersController.DeleteTrainer)

	return server
}

func SetGinMode(profile string) {
	if profile == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
