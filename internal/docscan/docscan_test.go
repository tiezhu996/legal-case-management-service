package docscan

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func writeTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	for _, name := range []string{"a.pdf", "b.pdf", "c.txt"} {
		if err := os.WriteFile(filepath.Join(root, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestDocscanCancel(t *testing.T) {
	root := writeTree(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := ScanDocuments(ctx, root)
	if !errors.Is(err, ErrScanCanceled) {
		t.Fatalf("expected ErrScanCanceled, got %v", err)
	}
}

func TestDocscanWalk(t *testing.T) {
	root := writeTree(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	w := NewWalker(ctx)
	_, err := w.Walk(root)
	if !errors.Is(err, ErrScanCanceled) {
		t.Fatalf("expected ErrScanCanceled from Walker, got %v", err)
	}
}

func TestDocscanGather(t *testing.T) {
	root := writeTree(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	w := NewWalker(ctx)
	var out []string
	err := w.CollectFiles(root, &out)
	if !errors.Is(err, context.Canceled) && !errors.Is(err, ErrScanCanceled) {
		t.Fatalf("expected cancellation error, got %v", err)
	}
}

func TestDocscanTotal(t *testing.T) {
	root := writeTree(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := ScanDocumentCount(ctx, root)
	if !errors.Is(err, ErrScanCanceled) {
		t.Fatalf("expected ErrScanCanceled, got %v", err)
	}
}
