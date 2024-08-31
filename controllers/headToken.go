package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// Verify Access Token godoc
// @Summary Verify the access token
// @Description Verify the access token using the token in Authorization header
// @Tags Auth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Security ApiKeyAuth|OAuth2Application
// @Success 200 {string} string "Successfully verified access token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/token [head]
func VerifyToken(c *gin.Context) {
	// expects the token to be verified in the middleware
	utils.SendSuccessResponse(c, 200, "Successfully verified access token", nil)
}
