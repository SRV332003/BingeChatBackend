package controllers

import (
	"HangAroundBackend/services/googleauth"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// GetGoogleLoginUri godoc
// @Summary Get google login uri
// @Description Get the google login uri to redirect the user to google login page
// @Tags GoogleAuth
// @Produce  json
// @Success 200 {string} string "Successfully generated google login uri"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/google [get]
func GetGoogleLoginUri(c *gin.Context) {

	uri, state := googleauth.GetURL()
	utils.SendSuccessResponse(c, 200, "Successfully generated google login uri", gin.H{
		"uri":   uri,
		"state": state,
	})
	return
}
