package retention

import "time"

// retentionYears 文档类型 -> 保管年限。
var retentionYears = map[string]int{
	"judgment": 30, "evidence": 5, "contract": 10, "complaint": 10, "defense": 10, "other": 3,
}

// RetentionYears 返回某文档类型的档案保管年限，未知类型按 3 年。
func RetentionYears(fileType string) int {
	if y, ok := retentionYears[fileType]; ok {
		return y
	}
	return 3
}

// RetentionExpiry 返回档案保管到期日。
func RetentionExpiry(fileType string, createdAt time.Time) time.Time {
	return createdAt.AddDate(RetentionYears(fileType), 0, 0)
}

// IsExpired 判断档案是否已过保管期限。
func IsExpired(fileType string, createdAt, now time.Time) bool {
	return now.After(RetentionExpiry(fileType, createdAt))
}
