package deadline

import (
	"errors"
	"fmt"
	"time"
)

// 期限计算相关哨兵错误。
var (
	ErrInvalidCaseType = errors.New("invalid case type")
	ErrMissingDate     = errors.New("missing accept date")
	ErrFutureDate      = errors.New("date must not be in the future")
)

// ComputeDeadlines 计算审理期限与上诉期限，先做输入校验。
func ComputeDeadlines(caseType string, acceptDate time.Time) (trial, appeal time.Time, err error) {
	if _, ok := trialMonths[caseType]; !ok {
		return time.Time{}, time.Time{}, ErrInvalidCaseType
	}
	if acceptDate.IsZero() {
		return time.Time{}, time.Time{}, ErrMissingDate
	}
	if acceptDate.After(time.Now()) {
		return time.Time{}, time.Time{}, ErrFutureDate
	}
	trial = TrialDeadline(caseType, acceptDate)
	appeal = AppealDeadline(trial)
	return trial, appeal, nil
}

// ValidateCaseType 校验案件类型是否支持期限计算。
func ValidateCaseType(caseType string) error {
	if _, ok := trialMonths[caseType]; !ok {
		return fmt.Errorf("validate case type %q: %w", caseType, ErrInvalidCaseType)
	}
	return nil
}

// ClassifyDeadlineError 将期限计算错误归类为 type / date / future / other。
func ClassifyDeadlineError(err error) string {
	if err == nil {
		return "none"
	}
	switch {
	case errors.Is(err, ErrInvalidCaseType):
		return "type"
	case errors.Is(err, ErrMissingDate):
		return "date"
	case errors.Is(err, ErrFutureDate):
		return "future"
	default:
		return "other"
	}
}
