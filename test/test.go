package test

import (
	"github.com/gin-gonic/gin"
)

// PingHandler responds with pong
// @Summary Test Ping example
// @Description Responds with pong from test package
// @Success 200 {string} string "pong"
// @Router /test/ping [get]
func PingHandler(c *gin.Context) {
	c.String(200, "pong from test")
}
