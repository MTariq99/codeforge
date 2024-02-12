package middlewares

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(int64)
		email := c.MustGet("email").(string)

		// Create a new Sentry event.
		event := sentry.NewEvent()
		event.Message = fmt.Sprintf("Request received from %v", email)
		event.Extra["user_id"] = userId
		event.Extra["email"] = email
		event.Extra["method"] = c.Request.Method
		event.Extra["path"] = c.Request.URL.Path
		sentry.CaptureEvent(event)

		// Log the event using logrus.
		logger.WithFields(logrus.Fields{
			"user_id": userId,
			"email":   email,
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
		}).Info("Request received")

		c.Next()
	}
}

func NonAuthLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.Request.Header.Get("Authorization")
		logger.WithFields(logrus.Fields{
			"clientIP": clientIP,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
		}).Info("NonAuth Request received")

		c.Next()
	}

}
