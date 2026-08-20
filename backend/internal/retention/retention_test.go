package retention

import (
	"testing"
	"time"
)

func base() time.Time { return time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC) }

func TestRetentionPut(t *testing.T) {
	idx := NewRetentionIndex(nil)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Put panicked on empty index: %v", r)
		}
	}()
	idx.Put("evidence", 5)
	if y, ok := idx.Years("evidence"); !ok || y != 5 {
		t.Fatalf("expected years 5, got %d ok=%v", y, ok)
	}
}

func TestRetentionYear(t *testing.T) {
	idx := NewRetentionIndex([]string{"judgment", "evidence"})
	if y, ok := idx.Years("judgment"); !ok || y != 30 {
		t.Fatalf("judgment should be 30, got %d ok=%v", y, ok)
	}
	if _, ok := idx.Years("contract"); ok {
		t.Fatal("contract should not be present")
	}
}

func TestRetentionExpiry(t *testing.T) {
	created := base()
	expiry := RetentionExpiry("evidence", created)
	if expiry.Year() != created.Year()+5 {
		t.Fatalf("evidence expiry year should be +5, got %v", expiry)
	}
}
