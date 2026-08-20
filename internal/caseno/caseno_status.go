package caseno

// CaseNoHTTPStatus 将案件编号错误归类映射为 HTTP 状态码：合法 200，格式 400，年份/校验码 422。
func CaseNoHTTPStatus(s string) int {
	switch CaseNoErrorCode(s) {
	case 0:
		return 200
	case 1:
		return 400
	case 2, 3:
		return 422
	default:
		return 400
	}
}
