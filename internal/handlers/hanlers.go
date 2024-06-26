package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	PingHandler(c *gin.Context)
}

// HealthChecker struct provides the handler for a health check endpoint.
// This is a starting point which allows each application to
// customize their health check strategy.
type HealthChecker struct{}

func New() Interface {
	return HealthChecker{}
}

// PingHandler returns a successful pong answer to all HTTP requests.
func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
