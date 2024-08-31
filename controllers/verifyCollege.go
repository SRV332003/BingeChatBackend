package controllers

import (
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
)

type VerCollegeRequest struct {
	Name        string `json:"name" binding:"required"`
	CollegeID   string `json:"college_id" binding:"required"`
	EmailFormat string `json:"emailFormat" binding:"required"`
}

// AddCollege godoc
// @Summary Add a new college
// @Description Add a new college
// @Tags College
// @Accept  json
// @Produce  json
// @Param name formData string true "CollegeName"
// @Param emailFormat formData string true "Email Format"
// @Router /api/v1/college [post]
func VerCollege(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		utils.SendErrorResponse(c, 403, "You are not authorized to add college")
		return
	}

	var req VerCollegeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	name := req.Name
	college_id := req.CollegeID
	emailFormat := req.EmailFormat

	if name == "" || college_id == "" || emailFormat == "" {
		utils.SendErrorResponse(c, 400, "Please provide all the required fields")
		return
	}

	// Check if college already exists
	exists, college, err := crud.CheckCollegeExists(emailFormat)

	if !exists || err != nil {
		utils.SendErrorResponse(c, 400, "College does not exist")
		return
	}

	college.Verified = true

	err = crud.UpdateCollegeStatus(&college)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error Updating College status")
		return
	}

	utils.SendSuccessResponse(c, 200, "College Verified successfully!!", college)
}
