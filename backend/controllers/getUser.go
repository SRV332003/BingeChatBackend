package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// AddCashController godoc
// @Summary Add cash to user account
// @Description Add cash to user account
// @Tags Auth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {string} string "Successfully fetched user"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
