package controllers

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/customauth"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/services/googleauth"
	"HangAroundBackend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type verifyAuthCodeReq struct {
	Code string `json:"code" binding:"required"`
}

// VerifyAuthCode godoc
// @Summary Verify the google auth code
// @Description Verify the google auth code and return the access token with the user details
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param code formData string true "Google Auth Code"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/google [post]
func VerifyAuthCode(c *gin.Context) {
	var req verifyAuthCodeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	code := req.Code

	if code == "" {
		utils.SendErrorResponse(c, 400, "Auth code cannot be empty")
		return
	}

	token, err := googleauth.Exchange(code, c)
	if err != nil {
		utils.SendErrorResponse(c, 401, "Invalid google code")
		return
	}

	userData, err := googleauth.GetUserInfo(token.AccessToken)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	res, userLogin, err := crud.CheckUserLoginExists(userData["email"].(string))
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	if res {
		if userLogin.Email != userData["email"].(string) {
			utils.SendErrorResponse(c, 401, "An unauthorized user is trying to login")
		}

		// User already exists
		access_token, err := customauth.GenerateAccessToken(userLogin.ID, userLogin.Email, userLogin.Name, userLogin.Role)
		if err != nil {
			utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
			return
		}

		refresh_token, err := customauth.GenerateRefreshToken(userLogin.ID)
		if err != nil {
			utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
			return
		}

		utils.SendSuccessResponse(c, 200, "User successfully logged in", gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		})
		return
	}

	collegeEmailFormat := strings.Split(userData["email"].(string), "@")

	exists, college, err := crud.CheckCollegeExists(collegeEmailFormat[1])
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	if !exists {
		utils.SendErrorResponse(c, 403, "Your college is not registered with us.")
		return
	}

	access_token, err := customauth.GenerateAccessToken(userLogin.ID, userLogin.Email, userLogin.Name, userLogin.Role)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}
	refresh_token, err := customauth.GenerateRefreshToken(userLogin.ID)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	// generate random password
	password := utils.GenerateRandomPassword(10)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	// Create new user
	userLogin = models.UserLogin{
		Email:             userData["email"].(string),
		Name:              userData["name"].(string),
		Role:              "user",
		Password:          string(hashedPassword),
		CollegeID:         college.ID,
		Verified:          true,
		VerificationToken: "",
		RefreshToken:      refresh_token,
	}

	err = crud.CreateUserLogin(&userLogin)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, 200, "User partially registered, please complete the registration", gin.H{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}
