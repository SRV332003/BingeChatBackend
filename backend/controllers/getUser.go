package controllers

import (
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Fetch user account information
// @Description Fetch user account information
// @Tags Auth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {string} string "Successfully fetched user"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {
	email := c.GetString("email")

	user, err := crud.GetUserInfoByEmail(email)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	utils.SendSuccessResponse(c, 200, "Successfully fetched user", user)

}
