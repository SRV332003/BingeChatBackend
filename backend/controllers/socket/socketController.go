package socket

import (
	"HangAroundBackend/logger"

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
