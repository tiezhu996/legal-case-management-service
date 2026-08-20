package util

import (
	"testing"

	"cylawcase/internal/constants"
)

func TestCaseStatusTextSuspended(t *testing.T) {
	if got := CaseStatusText(constants.CaseStatusSuspended); got != "已中止" {
		t.Errorf("CaseStatusText(suspended) = %q, want 已中止", got)
	}
	if got := CaseStatusText(constants.CaseStatusClosed); got != "已结案" {
		t.Errorf("CaseStatusText(closed) = %q, want 已结案", got)
	}
}
