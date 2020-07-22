package controllers

import (
	"net/http"

	"github.com/janfalih17/api_pulsa/utils"

	"github.com/janfalih17/api_pulsa/configs"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	user := User{}
	c.BindJSON(&user)
	db, err := configs.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var result struct {
		ID       int
		Username string
		Password string
	}
	db.Table("users").Select("id, username, password").Where("username = ?", user.Username).Scan(&result)
	check := utils.CheckHash(user.Password, result.Password)
	if check != true {
		c.JSON(http.StatusUnauthorized, "Gagal Username Atau Password Salah")
		return
	}
	token, err := utils.CreateToken(result.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Gagal Membuat Token User")
		return
	}
	response := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	c.JSON(http.StatusOK, response)
	return

}
