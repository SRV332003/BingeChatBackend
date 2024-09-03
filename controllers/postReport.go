package controllers

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"
	"HangAroundBackend/utils/validators"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateReportReq struct {
	ReportedMail string `json:"reportedEmail" binding:"required"`
	ReportedText string `json:"reportedText" binding:"required"`
}

// Verify Access Token godoc
// @Summary Verify the access token
// @Description Verify the access token using the token in Authorization header
// @Tags Auth
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Security ApiKeyAuth|OAuth2Application
// @Success 200 {string} string "Successfully verified access token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/token [head]
func CreateReport(c *gin.Context) {
	// expects the token to be verified in the middleware
	var req CreateReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Bad Request: Invalid Json format")
		return
	}

	reporterMail := c.GetString("email")
	reportedMail := req.ReportedMail
	text := strings.TrimSpace(req.ReportedText)

	err := validators.IsValidEmail(reportedMail)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(text) > 1000 {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Too long text lengths are not allowed")
	}

	reporter, err := crud.GetUserLoginByEmail(reporterMail)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotAcceptable, "Reporter user with mail: \""+reporterMail+"\" was not found")
	}

	reported, err := crud.GetUserLoginByEmail(reportedMail)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotAcceptable, "Reported user with mail: \""+reportedMail+"\" was not found")
	}

	report := &models.Report{
		ReporterID: reporter.ID,
		ReportedID: reported.ID,
		ReportText: text,
	}

	ControllerLogger.Info("Person '" + reportedMail + "' was reported by '" + reporterMail + "'")
	err = crud.CreateReport(report)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error adding Report")
		return
	}

	utils.SendSuccessResponse(c, 200, "Successfully submitted report", nil)
}
