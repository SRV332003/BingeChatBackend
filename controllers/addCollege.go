package controllers

import (
	"HangAroundBackend/logger"
	"HangAroundBackend/models"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ControllerLogger *zap.Logger

type AddCollegeRequest struct {
	Name        string `json:"name" binding:"required"`
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
func AddCollege(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		utils.SendErrorResponse(c, 403, "You are not authorized to add college")
		return
	}

	var req AddCollegeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}
	log.Println(req)
	name := req.Name
	emailFormat := req.EmailFormat

	if name == "" || emailFormat == "" {
		utils.SendErrorResponse(c, 400, "Please provide all the required fields")
		return
	}

	college := models.College{
		Name:        name,
		EmailFormat: emailFormat,
		Verified:    true,
	}

	// Check if college already exists
	exists, _, _ := crud.CheckCollegeExists(emailFormat)
	if exists {
		utils.SendErrorResponse(c, 400, "College already exists")
		return
	}

	err := crud.CreateCollege(&college)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error adding college")
		return
	}

	utils.SendSuccessResponse(c, 200, "College added successfully", college)
}

func init() {
	ControllerLogger = logger.GetLoggerWithName("Controller")
}
