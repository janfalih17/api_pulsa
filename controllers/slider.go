package controllers

import (
	"net/http"

	"github.com/janfalih17/api_pulsa/configs"
	"github.com/janfalih17/api_pulsa/models"

	"github.com/gin-gonic/gin"
)

func GetSlider(c *gin.Context) {
	db, err := configs.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	result := db.Model(&models.User{}).Related(&models.Status{})
	c.JSON(http.StatusOK, result)
	return
}
