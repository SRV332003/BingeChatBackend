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

	if authHeader == "" {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
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
		c.AbortWithStatus(401)
		return
	}

	username := claims.Username

	c.Set("username", username)

	c.Next()

}
