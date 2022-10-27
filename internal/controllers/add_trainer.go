package controllers

import (
	"github.com/faveroferreira/maromba-center/internal/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) AddTrainer(ctx *gin.Context) {
	body := models.Trainer{}

	// Ugly as f-
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Create(&body); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &body)
}
