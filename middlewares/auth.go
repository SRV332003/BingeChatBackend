package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	customauth "HangAroundBackend/services/customauth"
	"HangAroundBackend/utils"
)

func AuthMiddlware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := c.Query("token")

	if authHeader == "" && token == "" {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
	}

	if authHeader == "" {
		authHeader = "Bearer " + token
	}

	//Expecting format of Bearer token

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
	}

	tokeString := parts[1]

	claims, err := customauth.VerifyAccessToken(tokeString)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	c.Set("email", claims.Email)
	c.Set("role", claims.Role)
	c.Set("name", claims.Username)

	c.Next()

}
