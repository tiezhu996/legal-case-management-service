package caseno

import "errors"

// ClassifyCaseNo 将案件编号解析错误归为三类之一；编号合法时返回 nil。
func ClassifyCaseNo(s string) error {
	if _, err := ParseCaseNo(s); err == nil {
		return nil
	} else {
		switch {
		case errors.Is(err, ErrCaseNoFormat):
			return ErrCaseNoFormat
		case errors.Is(err, ErrCaseNoYear):
			return ErrCaseNoYear
		case errors.Is(err, ErrCaseNoCheck):
			return ErrCaseNoCheck
		default:
			return ErrCaseNoFormat
		}
	}
}

// CaseNoErrorCode 返回案件编号错误的分类码：0 合法，1 格式，2 年份，3 校验码。
func CaseNoErrorCode(s string) int {
	err := ClassifyCaseNo(s)
	switch {
	case err == nil:
		return 0
	case errors.Is(err, ErrCaseNoFormat):
		return 1
	case errors.Is(err, ErrCaseNoYear):
		return 2
	case errors.Is(err, ErrCaseNoCheck):
		return 3
	default:
		return 1
	}
}
