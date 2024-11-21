package rate_limiter

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	commonredis "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/database/redis"
	"strconv"
	"time"
)

type (
	// RateLimiter interface
	RateLimiter interface {
		Limit(ip string) error
	}

	// DefaultRateLimiter struct
	DefaultRateLimiter struct {
		redisClient *redis.Client
		limit       int
		period      time.Duration
	}
)

// NewDefaultRateLimiter creates a new rate limiter
func NewDefaultRateLimiter(redisClient *redis.Client, limit int, period time.Duration) *DefaultRateLimiter {
	return &DefaultRateLimiter{
		redisClient: redisClient,
		limit:       limit,
		period:      period,
	}
}

// Limit limits the rate of requests
func (d *DefaultRateLimiter) Limit(ip string) error {
	key := commonredis.GetKey(ip, RateLimiterPrefix)

	// Check the current rate limit
	value, err := d.redisClient.Get(context.Background(), key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	// Parse value
	var count int64
	if err == nil {
		count, _ = strconv.ParseInt(value, 10, 64)
	} else {
		// Set the initial value
		_, err = d.redisClient.Set(context.Background(), key, 1, d.period).Result()
		return err
	}

	// If the rate limit is exceeded, return an error
	if count >= int64(d.limit) {
		return TooManyRequestsError
	}

	// Increment the request count
	err = d.redisClient.Incr(context.Background(), key).Err()
	if err != nil {
		return err
	}

	// Set the expiration time
	err = d.redisClient.Expire(context.Background(), key, d.period).Err()
	if err != nil {
		return err
	}

	return nil
}
