package controllers

import (
	"github.com/faveroferreira/maromba-center/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetTrainer(ctx *gin.Context) {
    id := ctx.Param("id")

    var trainer models.Trainer

    if result := h.DB.First(&trainer, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &trainer)
}

func (h Handler) GetTrainers(ctx *gin.Context) {
    var trainers []models.Trainer

    if result := h.DB.Find(&trainers); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &trainers)
}