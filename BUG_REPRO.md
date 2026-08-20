# Bug 复现说明（lexcase-casestatus-statemachine-010）

## Bug 是什么
案件状态机新增 suspended 中止态后，转换表缺边、进行中判断漏算、状态文案漏配，跨包状态错位。

## 如何触发
在 `backend/` 目录下运行：

```
go test ./internal/service -run '^TestCanFlowSuspended$' -count=1
```

## 错误信息
中止态流转被拒绝、进行中判断和文案展示均不一致。
