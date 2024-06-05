package mid

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log information about incoming request
		slog.Info("Incoming request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		// Process the request
		c.Next()

		// Log information about outgoing response
		slog.Info("Outgoing response",
			slog.Int("status", c.Writer.Status()),
		)
	}
}
