package caseno

import (
	"errors"
	"testing"
)

func TestCasenoYearChain(t *testing.T) {
	c, err := ParseCaseNo("CY202600014")
	if err != nil || c.Year != 2026 || c.Seq != 1 || c.Check != 4 {
		t.Fatalf("valid parse failed: %+v err=%v", c, err)
	}
	if _, err := ParseCaseNo("CY199900014"); !errors.Is(err, ErrCaseNoYear) {
		t.Fatalf("expected ErrCaseNoYear in chain, got %v", err)
	}
}

func TestCasenoCheckChain(t *testing.T) {
	if _, err := ParseCaseNo("CY202600019"); !errors.Is(err, ErrCaseNoCheck) {
		t.Fatalf("expected ErrCaseNoCheck in chain, got %v", err)
	}
	if _, err := ParseCaseNo("bad"); !errors.Is(err, ErrCaseNoFormat) {
		t.Fatalf("expected ErrCaseNoFormat, got %v", err)
	}
}

func TestCasenoKind(t *testing.T) {
	if err := ClassifyCaseNo("CY199900014"); !errors.Is(err, ErrCaseNoYear) {
		t.Fatalf("expected year class, got %v", err)
	}
	if err := ClassifyCaseNo("CY202600019"); !errors.Is(err, ErrCaseNoCheck) {
		t.Fatalf("expected check class, got %v", err)
	}
	if err := ClassifyCaseNo("bad"); !errors.Is(err, ErrCaseNoFormat) {
		t.Fatalf("expected format class, got %v", err)
	}
}

func TestCasenoCode(t *testing.T) {
	if CaseNoErrorCode("CY202600014") != 0 {
		t.Fatal("valid case no should have code 0")
	}
	if CaseNoErrorCode("CY199900014") != 2 {
		t.Fatal("year error should have code 2")
	}
	if CaseNoErrorCode("CY202600019") != 3 {
		t.Fatal("checksum error should have code 3")
	}
	if CaseNoErrorCode("bad") != 1 {
		t.Fatal("format error should have code 1")
	}
}

func TestCasenoHTTP(t *testing.T) {
	if CaseNoHTTPStatus("CY202600014") != 200 {
		t.Fatal("valid case no should map to 200")
	}
	if CaseNoHTTPStatus("bad") != 400 {
		t.Fatal("format error should map to 400")
	}
	if CaseNoHTTPStatus("CY199900014") != 422 {
		t.Fatal("year error should map to 422")
	}
	if CaseNoHTTPStatus("CY202600019") != 422 {
		t.Fatal("checksum error should map to 422")
	}
}
