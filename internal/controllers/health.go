package controllers

import "github.com/gin-gonic/gin"

type HealthController struct {
}

func (c *HealthController) HealthCheck(context *gin.Context) {
	context.JSON(200, "tu tá é lascado")
}
