package controllers

import (
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name        string `json:"name" `
	Branch      string `json:"branch" `
	PassoutYear string `json:"passout_year" `
	Course      string `json:"course" `
	RollNo      string `json:"roll_no" `
}

// UpdateUser godoc
// @Summary Update user account information
// @Description Update user account information
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// Param Authorization header string true "Bearer token"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/{id} [put]
func UpdateUser(c *gin.Context) {

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	name := req.Name
	if name != "" {
		// update name

	}

	// userinfo := models.UserInfo{
	// 	Rollno:      req.RollNo,
	// 	Branch:      req.Branch,
	// 	PassoutYear: req.PassoutYear,
	// 	Course:      req.Course,
	// }

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
