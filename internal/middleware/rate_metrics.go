package middleware

import (
	"net/http"
	"strconv"
)

// Stats 返回当前限流器每个 IP 的剩余令牌快照。
func (rl *RateLimiter) Stats() map[string]int {
	out := make(map[string]int, len(rl.buckets))
	for ip, b := range rl.buckets {
		out[ip] = b.tokens
	}
	return out
}

// Reset 清空某个 IP 的令牌桶。
func (rl *RateLimiter) Reset(ip string) {
	delete(rl.buckets, ip)
}

// RateMetricsHandler 暴露限流统计的 HTTP 处理器。
func (rl *RateLimiter) RateMetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats := rl.Stats()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"buckets":` + strconv.Itoa(len(stats)) + `}`))
	}
}
