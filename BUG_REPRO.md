# Bug 复现说明（lexcase-upload-defer-006）

## Bug 是什么
批量上传在循环内 defer 关闭、命名返回值 defer 吞错、错误分支漏关闭，写失败时接口返回成功且句柄延迟释放。

## 如何触发
在 `backend/` 目录下运行：

```
go test ./internal/upload -run '^TestUploadMkdir$' -count=1
```

## 错误信息
目录创建失败时 SaveBatch 仍返回 nil 错误。
