package controllers

import (
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// GetCollege godoc
// @Summary Fetch College information
// @Description Fetch College information
// @Tags Auth
// @Produce  json
// @Success 200 {string} string "Successfully fetched College"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user [get]
func GetCollege(c *gin.Context) {

	colleges, err := crud.GetAllCollegesNames()
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	utils.SendSuccessResponse(c, 200, "Successfully fetched College", colleges)
}
