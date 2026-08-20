package deadline

import "time"

// trialMonths 案件类型 -> 审理期限（月）。
var trialMonths = map[string]int{
	"civil": 6, "criminal": 2, "administrative": 6, "commercial": 6, "labor": 3,
}

// AddWorkingDays 从基准日向后推 N 个工作日（跳过周末）。
func AddWorkingDays(base time.Time, days int) time.Time {
	d := base
	added := 0
	for added < days {
		d = d.AddDate(0, 0, 1)
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			added++
		}
	}
	return d
}

// TrialDeadline 返回该案件类型自受理日起的审理期限届满日。
func TrialDeadline(caseType string, acceptDate time.Time) time.Time {
	months, ok := trialMonths[caseType]
	if !ok {
		months = 6
	}
	return acceptDate.AddDate(0, months, 0)
}

// AppealDeadline 返回自判决日起 15 个工作日的上诉期届满日。
func AppealDeadline(judgmentDate time.Time) time.Time {
	return AddWorkingDays(judgmentDate, 15)
}
