package upload

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SaveBatch 批量将内存中的文件内容写入目标目录，返回保存的文件名列表。
func SaveBatch(destDir string, files map[string]io.ReadCloser) (saved []string, err error) {
	if err := os.MkdirAll(destDir, 0o755); err != nil {
		return nil, fmt.Errorf("mkdir: %w", err)
	}
	for name, r := range files {
		if e := writeOne(filepath.Join(destDir, name), r); e != nil {
			r.Close()
			return nil, fmt.Errorf("save %s: %w", name, e)
		}
		if e := r.Close(); e != nil {
			return nil, fmt.Errorf("close %s: %w", name, e)
		}
		saved = append(saved, name)
	}
	return saved, nil
}

func writeOne(dst string, r io.Reader) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, r); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}
