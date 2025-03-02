package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengatur header untuk mencegah MIME-sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Mengatur header untuk mencegah website dibungkus dalam frame
		c.Header("X-Frame-Options", "DENY")
		// Mengaktifkan proteksi XSS pada browser
		c.Header("X-XSS-Protection", "1; mode=block")
		// Mengatur Content-Security-Policy untuk membatasi sumber daya yang dapat dimuat
		c.Header("Content-Security-Policy", "default-src 'self'")
		// Mengatur Referrer-Policy agar referrer tidak dikirim secara berlebihan
		c.Header("Referrer-Policy", "no-referrer")
		// Jika menggunakan HTTPS, sebaiknya aktifkan Strict-Transport-Security
		// c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		// Lanjutkan eksekusi request
		c.Next()
	}
}

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set header CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Lanjutkan eksekusi request
		c.Next()
	}
}

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
		err := c.Errors.Errors()
		// Log the information
		logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"latency":   latency,
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"error":     err,
		}).Info("Handled request")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
