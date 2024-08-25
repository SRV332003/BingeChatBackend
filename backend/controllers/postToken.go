package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Accepts user credentials and returns refresh and access tokens
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Successfully logged in"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/login [post]
func CreateToken(c *gin.Context) {
	// TODO: Implement get google auth

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
