package lib

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	redisClient *redis.Client
	ctxTest     = context.Background()
)

func TestRateLimitCheck(t *testing.T) {
	count := 3
	redisClient := NewRedisClient(false)
	redisClient.Set(ctxTest, "127.0.0.1", count, 5)
	period, calls := time.Minute, 10
	ratelimit := NewRateLimit(redisClient, period, calls)
	currentCount, remaining, err := ratelimit.Check("127.0.0.1")

	assert.Equal(t, currentCount, count+1)
	assert.Equal(t, remaining, calls-count-1)
	assert.Equal(t, err, nil)
}

func TestRateLimitCheckWithRedisServerError(t *testing.T) {
	redisClient := NewRedisClient(true)

	period, calls := time.Minute, 10
	ratelimit := NewRateLimit(redisClient, period, calls)
	currentCount, remaining, err := ratelimit.Check("127.0.0.1")

	assert.Equal(t, currentCount, -1)
	assert.Equal(t, remaining, -1)
	assert.Equal(t, err.Error(), "Redis Server Error")
}

func NewRedisClient(withErr bool) *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	if withErr {
		mr.SetError("Redis Server Error")
	}

	return client
}
