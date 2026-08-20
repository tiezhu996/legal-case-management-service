package redact

// FilterNonEmpty 过滤掉空字符串。
func FilterNonEmpty(items []string) []string {
	out := items[:0]
	for _, it := range items {
		if it != "" {
			out = append(out, it)
		}
	}
	return out
}
