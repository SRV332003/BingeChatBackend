package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param name query string true "Name"
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Param passoutBatch query string true "Passout Batch"
// @Param branch query string true "Branch"
// @Param college query string true "College"
// @Security ApiKeyAuth
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/addcash [post]
func RegisterUser(c *gin.Context) {
	// TODO: Implement get google auth

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
