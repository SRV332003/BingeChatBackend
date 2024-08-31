package controllers

import (
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type verifyUserReq struct {
	Token string `json:"token" binding:"required"`
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param name json string true "Name"
// @Param email json string true "Email"
// @Param password json string true "Password"
// @Param collegeID json string true "College ID"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/addcash [post]
func VerifyUser(c *gin.Context) {

	var req verifyUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	token := req.Token

	if token == "" {
		utils.SendErrorResponse(c, 400, "Token is required")
		return
	}

	user, err := crud.GetUserByVerificationToken(token)
	if err != nil {
		utils.SendErrorResponse(c, 400, "Invalid token")
		return
	}

	if user.Verified {
		utils.SendErrorResponse(c, http.StatusExpectationFailed, "User already verified")
		return
	}

	user.Verified = true
	user.VerificationToken = "a"
	err = crud.UpdateUserLogin(user)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, 200, "User Verification Successful", nil)
}

//$2a$10$0K.dJFvzg91EjtSzZVKbo.SL3i9mgVqEptCxBPCud.6CwK6d/bZny
//$2a$10$gP8UirrbS/VMcfeSYykmtOvI6To.kQzgdHnhpLdrhR7Ja4YutmBP2
