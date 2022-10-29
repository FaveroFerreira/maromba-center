package controllers

import (
	"net/http"

	"github.com/faveroferreira/maromba-center/internal/models"
	"github.com/faveroferreira/maromba-center/internal/services"
	"github.com/gin-gonic/gin"
)

type TrainerController struct {
	Services *services.TrainerServices
}

func (tc TrainerController) CreateTrainer(ctx *gin.Context) {
	body := models.Trainer{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := tc.Services.CreateTrainer(&body); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result)
		return
	}

	ctx.JSON(http.StatusCreated, &body)
}

func (tc TrainerController) GetTrainerById(ctx *gin.Context) {
	id := ctx.Param("id")
	var trainer models.Trainer

	if err := tc.Services.GetTrainerById(&trainer, id); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &trainer)
}

func (tc TrainerController) GetAllTrainers(ctx *gin.Context) {
	var trainer_list []models.Trainer

	if err := tc.Services.GetAllTrainers(&trainer_list); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, &trainer_list)
}

func (tc TrainerController) UpdateTrainer(ctx *gin.Context) {
	id := ctx.Param("id")
	requestBody := models.Trainer{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updated models.Trainer

	if err := tc.Services.UpdateTrainer(&requestBody, &updated, id); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &updated)
}

func (tc TrainerController) DeleteTrainer(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := tc.Services.DeleteTrainer(id); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusOK)
}
