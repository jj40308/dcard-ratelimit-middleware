package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRemainingRequests(c *gin.Context) {
	remaining, _ := c.Get("remaining")
	c.JSON(http.StatusOK, map[string]interface{}{"status": "Success", "remaining_requests": remaining})

	return
}
