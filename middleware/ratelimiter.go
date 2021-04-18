package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jj40308/dcard-ratelimit-middleware/lib"
	"net/http"
	"strconv"
)

type RateLimiter struct {
	rateLimit *lib.RateLimit
}

func NewRateLimiter(rateLimit *lib.RateLimit) *RateLimiter {
	return &RateLimiter{
		rateLimit: rateLimit,
	}
}

func (r *RateLimiter) CheckRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		remaining, err := r.rateLimit.Check(ip)

		if remaining < 0 {
			if err != nil {
				sendErrorResponse(c, map[string]interface{}{"error": err})
				return
			}
			sendTooManyRequestResponse(c, map[string]interface{}{"error": "Too Many Requests."})
			return
		}

		c.Set("remaining", remaining)
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
	}
}

func sendErrorResponse(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, data)
}

func sendTooManyRequestResponse(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusTooManyRequests, data)
}
