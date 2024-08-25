package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// VerifyAuthCode godoc
// @Summary Verify the google auth code
// @Description Verify the google auth code and return the access token with the user details
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param code query string true "Google Auth Code"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/google [post]
func VerifyAuthCode(c *gin.Context) {
	// TODO: Implement get google auth

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
