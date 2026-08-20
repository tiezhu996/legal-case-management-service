package service

import (
	"testing"

	"cylawcase/internal/constants"
)

func TestCanFlowSuspended(t *testing.T) {
	cases := []struct {
		from, to string
		want     bool
	}{
		{constants.CaseStatusInvestigating, constants.CaseStatusSuspended, true},
		{constants.CaseStatusSuspended, constants.CaseStatusInvestigating, true},
		{constants.CaseStatusHearing, constants.CaseStatusSuspended, true},
		{constants.CaseStatusSuspended, constants.CaseStatusHearing, true},
	}
	for _, tc := range cases {
		if got := canFlow(tc.from, tc.to); got != tc.want {
			t.Errorf("canFlow(%s->%s) = %v, want %v", tc.from, tc.to, got, tc.want)
		}
	}
}

func TestIsCaseActiveSuspended(t *testing.T) {
	if !IsCaseActive(constants.CaseStatusSuspended) {
		t.Error("suspended should be active")
	}
	if !IsCaseActive(constants.CaseStatusFiled) {
		t.Error("filed should be active")
	}
	if IsCaseActive(constants.CaseStatusClosed) {
		t.Error("closed should not be active")
	}
	if IsCaseActive(constants.CaseStatusArchived) {
		t.Error("archived should not be active")
	}
}

func TestSuspendedStatusValid(t *testing.T) {
	if !constants.IsValidCaseStatus(constants.CaseStatusSuspended) {
		t.Error("suspended should be a valid case status")
	}
	if constants.IsValidCaseStatus("bogus") {
		t.Error("bogus should not be valid")
	}
}
