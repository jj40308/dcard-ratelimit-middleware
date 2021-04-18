package lib

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// RateLimit with redis client
type RateLimit struct {
	redisClient *redis.Client
	period      time.Duration
	calls       int
}

var ctx = context.Background()

// NewRateLimit is a factory for RateLimit
func NewRateLimit(redisClient *redis.Client, period time.Duration, calls int) *RateLimit {
	return &RateLimit{
		redisClient: redisClient,
		period:      period,
		calls:       calls,
	}
}

// Check the remaining requests
func (r *RateLimit) Check(ip string) (int, error) {
	rc := r.redisClient
	value, errGet := rc.Get(ctx, ip).Result()
	if errGet != nil && errGet != redis.Nil {
		return -1, errGet
	}
	count, _ := strconv.Atoi(value)

	count = count + 1
	remaining := r.calls - count

	if remaining == -1 {
		return remaining, nil
	}

	errSet := rc.Set(ctx, ip, count, r.period).Err()
	if errSet != nil {
		return -1, errSet
	}

	return remaining, nil
}
