package middleware

import (
	"net/http"
	"sync"
	"time"

	"cylawcase/internal/constants"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	tokens   int
	lastFill time.Time
}

// RateLimiter 基于 IP 的简易令牌桶限流。
type RateLimiter struct {
	mu       sync.Mutex
	perMin   int
	buckets  map[string]*bucket
	capacity int
}

// NewRateLimiter 构造限流器。
func NewRateLimiter(perMin int) *RateLimiter {
	if perMin <= 0 {
		perMin = 120
	}
	return &RateLimiter{perMin: perMin, buckets: make(map[string]*bucket), capacity: perMin}
}

// Allow 判断某 IP 当前是否允许通过，并扣减一个令牌。
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	b, ok := rl.buckets[ip]
	now := time.Now()
	if !ok {
		b = &bucket{tokens: rl.capacity, lastFill: now}
		rl.buckets[ip] = b
	}
	rl.mu.Unlock()
	elapsed := now.Sub(b.lastFill)
	b.tokens += int(elapsed.Minutes()) * rl.perMin
	if b.tokens > rl.capacity {
		b.tokens = rl.capacity
	}
	b.lastFill = now
	if b.tokens <= 0 {
		return false
	}
	b.tokens--
	return true
}

// Peek 返回某 IP 当前剩余令牌数，不扣减。
func (rl *RateLimiter) Peek(ip string) int {
	if b, ok := rl.buckets[ip]; ok {
		return b.tokens
	}
	return rl.capacity
}

// Limit 返回限流中间件。
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rl.Allow(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"code": constants.CodeTooManyRequests, "message": constants.MsgTooManyRequests, "data": nil})
			return
		}
		c.Next()
	}
}

// Snapshot 返回限流桶快照。
func (rl *RateLimiter) Snapshot() map[string]*bucket {
	return rl.buckets
}
