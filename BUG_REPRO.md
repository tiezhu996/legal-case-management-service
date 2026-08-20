# Bug 复现说明（lexcase-timeline-slice-004）

## Bug 是什么
时间线的过滤、合并、切片、去重原地复用底层数组或共享切片，污染调用方原切片内容。

## 如何触发
在 `backend/` 目录下运行：

```
go test ./internal/timeline -run '^TestTimelinePick$' -count=1
```

## 错误信息
过滤后再读原列表，原列表条目内容被覆盖。
