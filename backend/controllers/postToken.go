package controllers

import (
	"HangAroundBackend/services/customauth"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateTokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login
// @Description Accepts user credentials and returns refresh and access tokens
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "Successfully logged in"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/user/login [post]
func CreateToken(c *gin.Context) {
	var req CreateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	email := req.Email
	password := req.Password

	if email == "" {
		utils.SendErrorResponse(c, 400, "Email is required")
		return
	}
	if password == "" {
		utils.SendErrorResponse(c, 400, "Password is required")
		return
	}

	// validate email and password
	user, err := crud.GetUserLoginByEmail(email)
	if err != nil {
		utils.SendErrorResponse(c, 401, "Invalid email or password")
		return
	}

	if user == nil {
		utils.SendErrorResponse(c, 401, "Invalid email or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		utils.SendErrorResponse(c, 401, "Invalid email or password")
		return
	}

	if !user.Verified {
		utils.SendErrorResponse(c, 401, "User not verified, please check your email for verification link")
		return
	}

	// generate refresh and access tokens
	refreshToken, err := customauth.GenerateRefreshToken(user.ID)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}
	accessToken, err := customauth.GenerateAccessToken(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	user.RefreshToken = refreshToken
	err = crud.UpdateUserLogin(user)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, 200, "Successfully logged in", gin.H{
		"refresh_token": refreshToken,
		"access_token":  accessToken,
	})
}
