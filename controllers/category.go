package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {

	c.JSON(http.StatusOK, "result")
}
