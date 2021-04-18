package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jj40308/dcard-ratelimit-middleware/controllers"
	"github.com/jj40308/dcard-ratelimit-middleware/lib"
	"github.com/jj40308/dcard-ratelimit-middleware/middleware"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func initRedis() *redis.Client {
	// Init redis client
	redisAddr := os.Getenv("REDIS_ADDR")
	redisOptions := redis.Options{
		Network: "tcp",
		Addr:    redisAddr,
	}
	redisClient := redis.NewClient(&redisOptions)
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Panic("Redis connection initialization failed:", err)
	}

	return redisClient
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := initRouter()
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func initRouter() *gin.Engine {
	r := gin.Default()

	redisClient := initRedis()
	period, calls := time.Minute, 60
	rateLimit := lib.NewRateLimit(redisClient, period, calls)
	rateLimiter := middleware.NewRateLimiter(rateLimit)

	apiv1 := r.Group("/v1")
	apiv1.GET("/get-remaining-requests", rateLimiter.CheckRateLimit(), controllers.GetRemainingRequests)

	return r
}
