package deadline

import (
	"errors"
	"testing"
	"time"
)

func TestDeadlineValidate(t *testing.T) {
	err := ValidateCaseType("bogus")
	if !errors.Is(err, ErrInvalidCaseType) {
		t.Fatalf("expected ErrInvalidCaseType in chain, got %v", err)
	}
	if err := ValidateCaseType("civil"); err != nil {
		t.Fatalf("civil should be valid, got %v", err)
	}
}

func TestDeadlineType(t *testing.T) {
	_, _, err := ComputeDeadlines("bogus", time.Now())
	if !errors.Is(err, ErrInvalidCaseType) {
		t.Fatalf("expected ErrInvalidCaseType, got %v", err)
	}
}

func TestDeadlineDate(t *testing.T) {
	_, _, err := ComputeDeadlines("civil", time.Time{})
	if !errors.Is(err, ErrMissingDate) {
		t.Fatalf("expected ErrMissingDate, got %v", err)
	}
}

func TestDeadlineBucket(t *testing.T) {
	_, _, err := ComputeDeadlines("bogus", time.Now())
	if ClassifyDeadlineError(err) != "type" {
		t.Fatal("expected type classification")
	}
	_, _, err = ComputeDeadlines("civil", time.Time{})
	if ClassifyDeadlineError(err) != "date" {
		t.Fatal("expected date classification")
	}
}

func TestDeadlineAccept(t *testing.T) {
	if err := ValidateAcceptDate(time.Time{}); !errors.Is(err, ErrMissingDate) {
		t.Fatalf("expected ErrMissingDate in chain, got %v", err)
	}
	future := time.Now().Add(24 * time.Hour)
	if err := ValidateAcceptDate(future); !errors.Is(err, ErrFutureDate) {
		t.Fatalf("expected ErrFutureDate in chain, got %v", err)
	}
	if err := ValidateAcceptDate(time.Now().Add(-24 * time.Hour)); err != nil {
		t.Fatalf("past date should be valid, got %v", err)
	}
}
