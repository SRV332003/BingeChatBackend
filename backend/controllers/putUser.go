package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// AddCashController godoc
// @Summary Add cash to user account
// @Description Add cash to user account
// @Tags payment
// @Accept  json
// @Produce  json
// @Param amount body int true "Amount to add"
// @Security ApiKeyAuth
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/addcash [post]
func UpdateUser(c *gin.Context) {
	// TODO: Implement get google auth

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
