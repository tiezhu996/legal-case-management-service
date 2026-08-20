# Bug 复现说明（lexcase-ratelimit-race-001）

## Bug 是什么

基于 IP 的令牌桶限流器 `middleware.RateLimiter` 存在锁范围漂移：`Allow` 在释放互斥锁之后继续修改桶的 `tokens` / `lastFill`，`Stats`、`Reset`、`Peek`、`Snapshot` 对共享 map 与桶字段的访问没有加锁。并发请求时对同一块内存读写产生 data race，服务偶发崩溃。

## 如何触发

在 `backend/` 目录下运行并发定向测试（带 `-race`）：

```
go test -race ./internal/middleware -run '^TestRateLimiterConcurrentAllow$' -count=1
```

## 错误信息

```
WARNING: DATA RACE
Read at 0x00c00025f108 by goroutine 8:
  cylawcase/internal/middleware.(*RateLimiter).Allow()
      internal/middleware/rate_limiter.go:44
Write at 0x00c00025f100 by goroutine 11:
  cylawcase/internal/middleware.(*RateLimiter).Allow()
      internal/middleware/rate_limiter.go:45
```
