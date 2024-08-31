package socket

import (
	"HangAroundBackend/logger"
	"HangAroundBackend/services/db/crud"
	"HangAroundBackend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var SocketLogger *zap.Logger
var manager *Manager

// AddCashController godoc
// @Summary Add cash to user account
// @Description Add cash to user account
// @Tags payment
// @Accept  json
// @Produce  json
// @Param amount body int true "Amount to add"
// @Success 200 {string} string "Successfully added cash"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /payment/addcash [post]
func SocketController(c *gin.Context) {

	email := c.GetString("email")
	verified, err := crud.CheckUserVerified(email)
	if err != nil {
		SocketLogger.Error("Error in checking user verification", zap.Error(err))
		utils.SendErrorResponse(c, 500, "Internal Server Error")
		return
	}

	if !verified {
		SocketLogger.Warn("User not verified")
		utils.SendErrorResponse(c, http.StatusNotAcceptable, "User not verified, please verify from link sent to your email")
		return
	}

	manager.HandleConnections(c)
}

func init() {
	SocketLogger = logger.GetLoggerWithName("socket")
	manager = NewManager()
	go manager.RoomDispatcher()
}

func CloseSocket() {
	manager.Close()
}
