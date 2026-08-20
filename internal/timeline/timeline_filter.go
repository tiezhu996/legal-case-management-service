package timeline

// FilterEvents 保留指定 kind 的时间线条目。
// 返回新分配的切片，不会修改 events 的底层数组。
func FilterEvents(events []Event, kind string) []Event {
	out := make([]Event, 0, len(events))
	for _, e := range events {
		if e.Kind == kind {
			out = append(out, e)
		}
	}
	return out
}
