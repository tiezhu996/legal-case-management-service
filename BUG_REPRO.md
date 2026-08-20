# Bug 复现说明（lexcase-policy-nilmap-003）

## Bug 是什么
权限加载器在空规则集下未初始化内部 map，typed-nil provider 也绕过 nil 判断，调用写入方法时触发 nil map 写入 panic。

## 如何触发
在 `backend/` 目录下运行：

```
go test ./internal/policy -run '^TestPolicyLoaderAdd$' -count=1
```

## 错误信息
```
panic: assignment to entry in nil map [recovered, repanicked]
cylawcase/internal/policy.(*PolicyLoader).Add(...)
	internal/policy/loader.go:15
```
