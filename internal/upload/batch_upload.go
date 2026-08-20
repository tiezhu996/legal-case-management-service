package upload

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SaveBatch 批量将内存中的文件内容写入目标目录，返回保存的文件名列表。
// 任意文件写入失败即返回错误；无论成功与否，所有传入的 ReadCloser 都会被关闭，
// 关闭失败会作为错误返回（仅在没有写入错误时）。
func SaveBatch(destDir string, files map[string]io.ReadCloser) (saved []string, err error) {
	if len(files) == 0 {
		return nil, fmt.Errorf("no files to save")
	}
	// 保证任意返回路径下所有 ReadCloser 都被关闭，并收集首个关闭错误。
	defer func() {
		var closeErr error
		for name, r := range files {
			if e := r.Close(); e != nil && closeErr == nil {
				closeErr = fmt.Errorf("close %s: %w", name, e)
			}
		}
		// 写入错误优先于关闭错误。
		if err == nil {
			err = closeErr
		}
	}()
	if e := os.MkdirAll(destDir, 0o755); e != nil {
		return nil, fmt.Errorf("mkdir: %w", e)
	}
	for name, r := range files {
		if e := writeOne(filepath.Join(destDir, name), r); e != nil {
			return nil, fmt.Errorf("save %s: %w", name, e)
		}
		saved = append(saved, name)
	}
	return saved, nil
}

func writeOne(dst string, r io.Reader) (err error) {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	// 确保目标文件句柄在任意返回路径下都被关闭；写入成功时把关闭错误返回给调用方。
	defer func() {
		if cerr := out.Close(); err == nil {
			err = cerr
		}
	}()
	if _, e := io.Copy(out, r); e != nil {
		return fmt.Errorf("copy: %w", e)
	}
	return nil
}
