package redact

import "strings"

// RedactIDNumber 脱敏身份证号：保留前 4 后 4，中间以 * 填充。
func RedactIDNumber(s string) string {
	if len(s) < 8 {
		return strings.Repeat("*", len(s))
	}
	return s[:4] + strings.Repeat("*", len(s)-8) + s[len(s)-4:]
}

// RedactPhone 脱敏手机号：保留前 3 后 4。
func RedactPhone(s string) string {
	if len(s) < 7 {
		return strings.Repeat("*", len(s))
	}
	return s[:3] + "****" + s[len(s)-4:]
}

// RedactName 脱敏姓名：仅保留首尾字符。
func RedactName(s string) string {
	r := []rune(s)
	switch {
	case len(r) == 0:
		return ""
	case len(r) == 1:
		return s
	case len(r) == 2:
		return string(r[0]) + "*"
	default:
		return string(r[0]) + strings.Repeat("*", len(r)-2) + string(r[len(r)-1])
	}
}
