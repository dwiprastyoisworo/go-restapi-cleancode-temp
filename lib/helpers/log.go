package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func CustomLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get request start time
		start := time.Now()

		// Process request
		c.Next()

		// Get request latency
		latency := time.Since(start)

		// Get some useful information
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log the information
		logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"latency":   latency,
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
		}).Info("Handled request")
	}
}
