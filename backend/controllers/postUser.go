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
// @Param name formData string true "Name"
// @Param rollNo formData string true "Roll No"
// @Param branch formData string true "Branch"
// @Param passoutBatch formData string true "Passout Batch"
// @Param course formData string true "Course"
// @Param college_id formData string true "College ID"
// @Param email formData string true "Email"
// @Param password formData string true "Password" 	
// @Security ApiKeyAuth
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/addcash [post]
func RegisterUser(c *gin.Context) {

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
