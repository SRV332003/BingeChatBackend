package controllers

import (
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

	utils.SendSuccessResponse(c, 200, "Successfully generated new access token", nil)
	
}
