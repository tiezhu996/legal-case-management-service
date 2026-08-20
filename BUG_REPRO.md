# Bug 复现说明（lexcase-caseno-errorchain-002）

## Bug 是什么
案号解析函数用 %v 包装哨兵错误，导致错误链断裂、errors.Is 失效，年份越界和校验码不符都被归类成格式错误，HTTP 状态码映射也随之误判。

## 如何触发
在 `backend/` 目录下运行：

```
go test ./internal/caseno -run '^TestCasenoYearChain$' -count=1
```

## 错误信息
`errors.Is(err, ErrCaseNoYear)` 返回 false，年份越界被归为格式错误。
