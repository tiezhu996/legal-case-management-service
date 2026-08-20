package upload

import (
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

func TestUploadMkdir(t *testing.T) {
	root := t.TempDir()
	blocker := filepath.Join(root, "notadir")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	destDir := filepath.Join(blocker, "sub")
	files := map[string]io.ReadCloser{
		"a.pdf": &closeTracker{name: "a.pdf", reader: strings.NewReader("x")},
	}
	if _, err := SaveBatch(destDir, files); err == nil {
		t.Fatal("expected error when mkdir fails")
	}
}

func TestSaveBatchClosesEveryFile(t *testing.T) {
	destDir := filepath.Join(t.TempDir(), "uploads")
	trackers := map[string]*closeTracker{
		"a.pdf": {name: "a.pdf", reader: strings.NewReader("content-a")},
		"b.pdf": {name: "b.pdf", reader: strings.NewReader("content-b")},
		"c.pdf": {name: "c.pdf", reader: strings.NewReader("content-c")},
	}
	files := map[string]io.ReadCloser{}
	for name, tr := range trackers {
		files[name] = tr
	}
	if _, err := SaveBatch(destDir, files); err != nil {
		t.Fatalf("SaveBatch failed: %v", err)
	}
	for name, tr := range trackers {
		if tr.closed != 1 {
			t.Fatalf("file %s closed %d times, want 1", name, tr.closed)
		}
	}
}

func TestSaveBatchWritesAllFiles(t *testing.T) {
	destDir := filepath.Join(t.TempDir(), "uploads")
	files := map[string]io.ReadCloser{
		"a.pdf": &closeTracker{name: "a.pdf", reader: strings.NewReader("content-a")},
	}
	saved, err := SaveBatch(destDir, files)
	if err != nil {
		t.Fatalf("SaveBatch failed: %v", err)
	}
	if len(saved) != 1 || saved[0] != "a.pdf" {
		t.Fatalf("unexpected saved: %v", saved)
	}
}
