package controllers

import (
	"HangAroundBackend/services/customauth"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// ReCreateToken godoc
// @Summary Refresh the access token
// @Description Renew the access token using the refresh token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param refresh_token query string true "Refresh token"
// @Success 200 {string} string "Successfully generated new access token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/token [get]
func ReCreateToken(c *gin.Context) {
	token := c.Query("refresh_token")

	if token == "" {
		utils.SendErrorResponse(c, 400, "Refresh token is required")
		return
	}

	claims, err := customauth.VerifyRefreshToken(token)
	if err != nil {
		utils.SendErrorResponse(c, 401, "Invalid refresh token")
		return
	}

	// generate new access token
	user, err := crud.GetUserLoginById(claims.ID)

	accessToken, err := customauth.GenerateAccessToken(user.ID, user.Email, user.Name, user.Role)

	utils.SendSuccessResponse(c, 200, "Successfully generated new access token", gin.H{
		"access_token": accessToken,
	})

}
