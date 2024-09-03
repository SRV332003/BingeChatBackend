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
// @Param code body string true "Google Auth Code"
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

	res, userLogin, _ := crud.CheckUserLoginExists(userData["email"].(string))

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
			"name":          userLogin.Name,
			"email":         userLogin.Email,
		})
		return
	}

	splitMail := strings.Split(userData["email"].(string), "@")

	exists, college, err := crud.CheckCollegeExists(splitMail[1])
	if !exists || err != nil {
		utils.SendErrorResponse(c, 403, "Your college is not registered with us.")
		return
	}

	// generate random password
	password := utils.GenerateRandomPassword(10)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	if userData["name"] == nil {
		userData["name"] = splitMail[0]
	}

	userInfo := models.UserInfo{}
	err = crud.CreateUserInfo(&userInfo)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Error creating User Information entry")
		return
	}

	// Create new user
	userLogin = models.UserLogin{
		UserInfo:          userInfo.ID,
		Email:             userData["email"].(string),
		Name:              userData["name"].(string),
		Role:              "user",
		Password:          string(hashedPassword),
		CollegeID:         college.ID,
		Verified:          true,
		VerificationToken: "",
	}

	err = crud.CreateUserLogin(&userLogin)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	// generate refresh and access tokens
	refreshToken, err := customauth.GenerateRefreshToken(userLogin.ID)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}
	accessToken, err := customauth.GenerateAccessToken(userLogin.ID, userLogin.Email, userLogin.Name, userLogin.Role)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	userLogin.RefreshToken = refreshToken
	err = crud.UpdateUserLogin(&userLogin)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	ControllerLogger.Info("Created user via Google Auth: " + userLogin.Email)
	utils.SendSuccessResponse(c, 200, "User registered Successfully", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"name":          userLogin.Name,
		"email":         userLogin.Email,
	})
}
