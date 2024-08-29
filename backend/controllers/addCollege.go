package controllers

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

// AddCollege godoc
// @Summary Add a new college
// @Description Add a new college
// @Tags College
// @Accept  json
// @Produce  json
// @Param name formData string true "CollegeName"
// @Param emailFormat formData string true "Email Format"
// @Router /api/v1/college [post]
func AddCollege(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		utils.SendErrorResponse(c, 403, "You are not authorized to add college")
		return
	}

	name := c.PostForm("name")
	college_id := c.PostForm("college_id")
	emailFormat := c.PostForm("emailFormat")

	if name == "" || college_id == "" || emailFormat == "" {
		utils.SendErrorResponse(c, 400, "Please provide all the required fields")
		return
	}

	college := models.College{
		Name:        name,
		EmailFormat: emailFormat,
		Verified:    false,
	}

	// Check if college already exists
	exists, _, err := crud.CheckCollegeExists(emailFormat)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error checking college")
		return
	}

	if exists {
		utils.SendErrorResponse(c, 400, "College already exists")
		return
	}

	err = crud.CreateCollege(&college)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error adding college")
		return
	}

	utils.SendSuccessResponse(c, 200, "College added successfully", college)
}
