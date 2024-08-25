package main

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"HangAroundBackend/services/db"
)

func shutdown(srv *http.Server, logger *zap.Logger) func(reason interface{}) {
	// handle graceful shutdown
	return func(reason interface{}) {
		logger.Sugar().Infof("Shutting down server: %v\n", reason)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.Disconnect()
		logger.Info("Disconnecting from DB")
		if err := srv.Shutdown(ctx); err != nil {
			logger.Info("Server Shutdown:" + err.Error())
		}
	}
}
