package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary Update user account information
// @Description Update user account information
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param body body UpdateUserRequest true "User information"
// Param Authorization header string true "Bearer token"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router ]
func UpdateUser(c *gin.Context) {
	// TODO: Implement get google auth
	
	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
