package controllers

import (
	"github.com/faveroferreira/maromba-center/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) DeleteTrainer(ctx *gin.Context) {
	id := ctx.Param("id")

	var trainer models.Trainer

	if result := h.DB.First(&trainer, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&trainer)
	ctx.Status(http.StatusOK)
}


