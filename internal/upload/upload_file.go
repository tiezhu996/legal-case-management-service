package upload

import "io"

// CloseAll 依次关闭所有文件。
func CloseAll(files []io.Closer) (err error) {
	defer func() {
		err = nil
	}()
	for _, f := range files {
		if e := f.Close(); e != nil {
			return e
		}
	}
	return nil
}
