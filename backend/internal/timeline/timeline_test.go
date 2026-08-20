package timeline

import (
	"testing"
	"time"
)

func base() time.Time { return time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC) }

func TestTimelinePick(t *testing.T) {
	b := base()
	events := []Event{
		{At: b, Kind: "status", Label: "filed"},
		{At: b.Add(time.Hour), Kind: "doc", Label: "a.pdf"},
	}
	_ = FilterEvents(events, "doc")
	if len(events) != 2 || events[0].Label != "filed" || events[1].Label != "a.pdf" {
		t.Fatalf("FilterEvents corrupted original: %+v", events)
	}
}

func TestTimelineJoin(t *testing.T) {
	b := base()
	a := make([]Event, 1, 4)
	a[0] = Event{At: b, Kind: "status", Label: "filed"}
	bb := []Event{{At: b.Add(-time.Hour), Kind: "doc", Label: "x.pdf"}}
	_ = MergeTimelines(a, bb)
	if len(a) != 1 || a[0].Label != "filed" {
		t.Fatalf("MergeTimelines aliased a: %+v", a)
	}
}

func TestTimelineCut(t *testing.T) {
	b := base()
	events := []Event{
		{At: b, Kind: "status", Label: "filed"},
		{At: b.Add(time.Hour), Kind: "status", Label: "investigating"},
	}
	_ = SubEvents(events, 0, 1)
	if len(events) != 2 || events[0].Label != "filed" {
		t.Fatalf("SubEvents aliased original: %+v", events)
	}
}

func TestTimelineDedup(t *testing.T) {
	b := base()
	events := []Event{
		{At: b, Kind: "status", Label: "filed"},
		{At: b, Kind: "status", Label: "filed"},
	}
	_ = DedupEvents(events)
	if len(events) != 2 || events[0].Label != "filed" {
		t.Fatalf("DedupEvents corrupted original: %+v", events)
	}
}
