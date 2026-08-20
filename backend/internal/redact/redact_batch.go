package redact

// RedactBatch 批量脱敏身份证号，返回全新切片，不影响原切片。
func RedactBatch(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, RedactIDNumber(it))
	}
	return out
}

// RedactAll 同时批量脱敏身份证号与手机号。
func RedactAll(idNumbers, phones []string) ([]string, []string) {
	return RedactBatch(idNumbers), RedactBatch(phones)
}

// RedactPhonesBatch 批量脱敏手机号，返回全新切片。
func RedactPhonesBatch(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, RedactPhone(it))
	}
	return out
}

// RedactNamesBatch 批量脱敏姓名，返回全新切片。
func RedactNamesBatch(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, RedactName(it))
	}
	return out
}
