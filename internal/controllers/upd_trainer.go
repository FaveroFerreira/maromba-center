package controllers

import (
	"github.com/faveroferreira/maromba-center/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) UpdateTrainer(ctx *gin.Context) {
    id := ctx.Param("id")
    body := models.Trainer{}

    if err := ctx.BindJSON(&body); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

	var trainer models.Trainer
    if result := h.DB.First(&trainer, id); result.Error != nil {
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }
    trainer.Name = body.Name
    trainer.Price = body.Price
    trainer.HatesGO = body.HatesGO

    h.DB.Save(&trainer)
    ctx.JSON(http.StatusOK, &trainer)
}