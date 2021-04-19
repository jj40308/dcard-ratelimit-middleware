package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRemainingRequests(c *gin.Context) {
	count, _ := c.Get("count")
	remaining, _ := c.Get("remaining")
	c.JSON(http.StatusOK, map[string]interface{}{"status": "Success", "count": count, "remaining_requests": remaining})

	return
}
