package controllers

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/customauth"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/services/mail"
	"HangAroundBackend/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CollegeID string `json:"collegeId" binding:"required"`
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
func RegisterUser(c *gin.Context) {

	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, 400, "Bad Request: "+err.Error())
		return
	}

	name := req.Name
	email := req.Email
	password := req.Password
	collegeId, err := strconv.Atoi(req.CollegeID)
	if err != nil {
		utils.SendErrorResponse(c, 400, "Invalid college ID")
		return
	}

	if name == "" || email == "" || password == "" || collegeId == 0 {
		utils.SendErrorResponse(c, 400, "All fields are required")
		return
	}

	// validate email and password

	// check if user exists
	_, err = crud.GetUserLoginByEmail(email)
	if err == nil {
		utils.SendErrorResponse(c, 401, "User already exists")
		return
	}

	// verify college
	college, err := crud.GetCollegeById(uint(collegeId))
	if err != nil {
		utils.SendErrorResponse(c, 400, "Invalid college ID")
		return
	}

	if !college.Verified {
		utils.SendErrorResponse(c, 400, "College not verified")
		return
	}

	if college.EmailFormat != strings.Split(email, "@")[1] {
		utils.SendErrorResponse(c, 400, "College email format does not match user email")
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	verifyToken, err := bcrypt.GenerateFromPassword([]byte(email+password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	// hashedPassword := password
	user := models.UserLogin{
		Name:              name,
		Email:             email,
		Password:          string(hashedPassword),
		CollegeID:         uint(collegeId),
		Role:              "user",
		VerificationToken: string(verifyToken),
		Verified:          false,
	}

	err = mail.SendVerificationMail([]string{email}, name, string(verifyToken))
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	err = crud.CreateUserLogin(&user)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
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
	err = crud.UpdateUserLogin(&user)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, 200, "User registered successfully, please verify your email", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"name":          user.Name,
	})
}

//$2a$10$0K.dJFvzg91EjtSzZVKbo.SL3i9mgVqEptCxBPCud.6CwK6d/bZny
//$2a$10$gP8UirrbS/VMcfeSYykmtOvI6To.kQzgdHnhpLdrhR7Ja4YutmBP2
