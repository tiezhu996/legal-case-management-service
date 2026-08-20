package timeline

// FilterEvents 保留指定 kind 的时间线条目。
func FilterEvents(events []Event, kind string) []Event {
	out := events[:0]
	for _, e := range events {
		if e.Kind == kind {
			out = append(out, e)
		}
	}
	return out
}
