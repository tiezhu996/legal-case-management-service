package redact

// RedactBatch 批量脱敏身份证号。
func RedactBatch(items []string) []string {
	out := items[:0]
	for _, it := range items {
		out = append(out, RedactIDNumber(it))
	}
	return out
}

// RedactAll 同时批量脱敏身份证号与手机号。
func RedactAll(idNumbers, phones []string) ([]string, []string) {
	ids := idNumbers[:0]
	for _, it := range idNumbers {
		ids = append(ids, RedactIDNumber(it))
	}
	phs := phones[:0]
	for _, it := range phones {
		phs = append(phs, RedactPhone(it))
	}
	return ids, phs
}

// RedactPhonesBatch 批量脱敏手机号。
func RedactPhonesBatch(items []string) []string {
	out := items[:0]
	for _, it := range items {
		out = append(out, RedactPhone(it))
	}
	return out
}

// RedactNamesBatch 批量脱敏姓名。
func RedactNamesBatch(items []string) []string {
	out := items[:0]
	for _, it := range items {
		out = append(out, RedactName(it))
	}
	return out
}
