package docscan

import (
	"context"
	"errors"
	"io/fs"
	"path/filepath"
)

// ErrScanCanceled 表示扫描因 context 取消/超时被中止。
var ErrScanCanceled = errors.New("scan canceled")

// ScanDocuments 递归扫描目录，收集所有非目录文件路径；支持 context 取消与超时。
func ScanDocuments(ctx context.Context, root string) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	out := make([]string, 0, 64)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if e := ctx.Err(); e != nil {
			return e
		}
		out = append(out, path)
		return nil
	})
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return nil, ErrScanCanceled
		}
		return nil, err
	}
	return out, nil
}
