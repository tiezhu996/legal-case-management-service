package docscan

import (
	"context"
	"io/fs"
	"path/filepath"
)

// Walker 持有 context 的目录遍历器，供长目录扫描复用。
type Walker struct {
	ctx context.Context
}

// NewWalker 构造遍历器，ctx 为空时退回 Background。
func NewWalker(ctx context.Context) *Walker {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Walker{ctx: ctx}
}

// Walk 遍历目录并返回文件路径列表。
func (w *Walker) Walk(root string) ([]string, error) {
	return ScanDocuments(w.ctx, root)
}

// CollectFiles 将目录下所有文件路径追加到 out。
func (w *Walker) CollectFiles(root string, out *[]string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if e := w.ctx.Err(); e != nil {
			return e
		}
		*out = append(*out, path)
		return nil
	})
}
