package controllers

import (
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

func IsUserVerified(c *gin.Context) {
	email := c.GetString("email")
	verified, err := crud.CheckUserVerified(email)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	if !verified {
		utils.SendErrorResponse(c, 417, "User not verified, please verify from link sent to your email")
		return
	}

	utils.SendSuccessResponse(c, 200, "User verified", nil)
}
