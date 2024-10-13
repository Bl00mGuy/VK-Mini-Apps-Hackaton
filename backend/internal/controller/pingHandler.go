package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
