package controllers

import (
	"HangAroundBackend/models"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CollegeID uint   `json:"collegeId" binding:"required"`
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
	collegeId := req.CollegeID

	if name == "" || email == "" || password == "" || collegeId == 0 {
		utils.SendErrorResponse(c, 400, "All fields are required")
		return
	}

	// validate email and password

	// check if user exists
	_, err := crud.GetUserLoginByEmail(email)
	if err == nil {
		utils.SendErrorResponse(c, 401, "User already exists")
		return
	}

	// verify college
	college, err := crud.GetCollegeById(collegeId)
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

	user := models.UserLogin{
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		CollegeID: collegeId,
		Role:      "user",
	}

	err = crud.CreateUserLogin(&user)
	if err != nil {
		utils.SendErrorResponse(c, 500, "Internal Server Error: "+err.Error())
		return
	}

	utils.SendSuccessResponse(c, 200, "Successfully added cash", nil)
}
