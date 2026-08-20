package upload

import "io"

// CloseAll 依次关闭所有文件，返回首个关闭错误，并保证全部文件都被尝试关闭。
func CloseAll(files []io.Closer) (err error) {
	for _, f := range files {
		if e := f.Close(); e != nil && err == nil {
			err = e
		}
	}
	return err
}
