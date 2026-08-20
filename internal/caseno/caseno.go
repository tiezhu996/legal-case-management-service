package caseno

import (
	"fmt"
	"strconv"
	"strings"
)

// CaseNo 案件编号解析结果。
type CaseNo struct {
	Prefix string
	Year   int
	Seq    int
	Check  int
}

// 案件编号相关哨兵错误。
var (
	ErrCaseNoFormat = fmt.Errorf("case number format invalid")
	ErrCaseNoYear   = fmt.Errorf("case number year invalid")
	ErrCaseNoCheck  = fmt.Errorf("case number checksum invalid")
)

// ComputeCheck 按加权和计算一位校验码：4 位序号各位 × 位权（1 起）之和模 10。
func ComputeCheck(seq int) int {
	digits := fmt.Sprintf("%04d", seq)
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i]-'0') * (i + 1)
	}
	return sum % 10
}

// FormatCaseNo 生成 "CY<年份><4位序号><校验码>" 形式的案件编号。
func FormatCaseNo(year, seq int) string {
	return fmt.Sprintf("CY%d%04d%d", year, seq, ComputeCheck(seq))
}

// ParseCaseNo 严格解析案件编号并校验校验码。
func ParseCaseNo(s string) (CaseNo, error) {
	if len(s) != 11 || !strings.HasPrefix(s, "CY") {
		return CaseNo{}, ErrCaseNoFormat
	}
	year, err := strconv.Atoi(s[2:6])
	if err != nil {
		return CaseNo{}, fmt.Errorf("parse year: %w", err)
	}
	if year < 2000 || year > 2999 {
		return CaseNo{}, fmt.Errorf("case year %d out of range: %w", year, ErrCaseNoYear)
	}
	seq, err := strconv.Atoi(s[6:10])
	if err != nil {
		return CaseNo{}, fmt.Errorf("parse seq: %w", err)
	}
	check, err := strconv.Atoi(s[10:11])
	if err != nil {
		return CaseNo{}, fmt.Errorf("parse check: %w", err)
	}
	if ComputeCheck(seq) != check {
		return CaseNo{}, fmt.Errorf("case %s checksum mismatch: %w", s, ErrCaseNoCheck)
	}
	return CaseNo{Prefix: "CY", Year: year, Seq: seq, Check: check}, nil
}

// ValidateCaseNo 判断案件编号是否合法。
func ValidateCaseNo(s string) bool {
	_, err := ParseCaseNo(s)
	return err == nil
}
