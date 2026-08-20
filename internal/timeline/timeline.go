package timeline

import (
	"sort"
	"time"
)

// StatusChange 案件状态变更记录。
type StatusChange struct {
	Status string
	At     time.Time
}

// DocEvent 案件文档事件。
type DocEvent struct {
	Title string
	At    time.Time
}

// Event 时间线条目。
type Event struct {
	At    time.Time
	Kind  string
	Label string
}

// BuildTimeline 将状态变更与文档事件合并为按时间升序的完整时间线。
func BuildTimeline(changes []StatusChange, docs []DocEvent) []Event {
	events := make([]Event, 0, len(changes)+len(docs))
	for _, c := range changes {
		events = append(events, Event{At: c.At, Kind: "status", Label: c.Status})
	}
	for _, d := range docs {
		events = append(events, Event{At: d.At, Kind: "doc", Label: d.Title})
	}
	sort.SliceStable(events, func(i, j int) bool { return events[i].At.Before(events[j].At) })
	return events
}

// MergeTimelines 合并两条已排序时间线并保持全局有序。
func MergeTimelines(a, b []Event) []Event {
	out := make([]Event, 0, len(a)+len(b))
	out = append(out, a...)
	out = append(out, b...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

// SubEvents 返回时间线 [from, to) 区间的副本。
func SubEvents(events []Event, from, to int) []Event {
	if from < 0 {
		from = 0
	}
	if to > len(events) {
		to = len(events)
	}
	if from > to {
		from = to
	}
	out := make([]Event, to-from)
	copy(out, events[from:to])
	return out
}

// DedupEvents 去除相邻重复条目，返回全新切片。
func DedupEvents(events []Event) []Event {
	out := make([]Event, 0, len(events))
	for _, e := range events {
		if len(out) > 0 && out[len(out)-1] == e {
			continue
		}
		out = append(out, e)
	}
	return out
}
