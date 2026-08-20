package upload

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type closeTracker struct {
	name   string
	closed int
	reader io.Reader
}

func (c *closeTracker) Read(p []byte) (int, error) { return c.reader.Read(p) }
func (c *closeTracker) Close() error               { c.closed++; return nil }

type errCloseTracker struct{ reader io.Reader }

func (e *errCloseTracker) Read(p []byte) (int, error) { return e.reader.Read(p) }
func (e *errCloseTracker) Close() error               { return errors.New("close boom") }

type failingReader struct{}

func (f *failingReader) Read(p []byte) (int, error) { return 0, errors.New("read boom") }

func TestUploadMkdir(t *testing.T) {
	root := t.TempDir()
	blocker := filepath.Join(root, "notadir")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	destDir := filepath.Join(blocker, "sub")
	files := map[string]io.ReadCloser{
		"a.pdf": &closeTracker{reader: strings.NewReader("x")},
	}
	if _, err := SaveBatch(destDir, files); err == nil {
		t.Fatal("expected error when mkdir fails")
	}
}

func TestUploadCloseFail(t *testing.T) {
	destDir := filepath.Join(t.TempDir(), "uploads")
	files := map[string]io.ReadCloser{
		"a.pdf": &errCloseTracker{reader: strings.NewReader("x")},
	}
	if _, err := SaveBatch(destDir, files); err == nil {
		t.Fatal("expected close error to propagate")
	}
}

func TestUploadCopyFail(t *testing.T) {
	destDir := filepath.Join(t.TempDir(), "uploads")
	files := map[string]io.ReadCloser{
		"a.pdf": &closeTracker{reader: &failingReader{}},
	}
	if _, err := SaveBatch(destDir, files); err == nil {
		t.Fatal("expected copy error to propagate")
	}
}

func TestUploadNoFile(t *testing.T) {
	destDir := filepath.Join(t.TempDir(), "uploads")
	if _, err := SaveBatch(destDir, map[string]io.ReadCloser{}); err == nil {
		t.Fatal("expected error for empty file set")
	}
}

func TestUploadCloseAll(t *testing.T) {
	var files []io.Closer
	files = append(files, &errCloseTracker{reader: strings.NewReader("x")})
	if err := CloseAll(files); err == nil {
		t.Fatal("expected close error to propagate")
	}
}
