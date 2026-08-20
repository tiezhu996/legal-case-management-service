package redact

import "testing"

func TestRedactIDs(t *testing.T) {
	items := []string{"110101199001011234", "110101199001015678"}
	orig := append([]string(nil), items...)
	_ = RedactBatch(items)
	if len(items) != 2 || items[0] != orig[0] || items[1] != orig[1] {
		t.Fatalf("RedactBatch corrupted original: %+v", items)
	}
}

func TestRedactPhone(t *testing.T) {
	items := []string{"13800000001", "13800000002"}
	orig := append([]string(nil), items...)
	_ = RedactPhonesBatch(items)
	if len(items) != 2 || items[0] != orig[0] || items[1] != orig[1] {
		t.Fatalf("RedactPhonesBatch corrupted original: %+v", items)
	}
}

func TestRedactName(t *testing.T) {
	items := []string{"张三", "李四"}
	orig := append([]string(nil), items...)
	_ = RedactNamesBatch(items)
	if len(items) != 2 || items[0] != orig[0] || items[1] != orig[1] {
		t.Fatalf("RedactNamesBatch corrupted original: %+v", items)
	}
}

func TestRedactBoth(t *testing.T) {
	ids := []string{"110101199001011234"}
	phones := []string{"13800000001"}
	idOrig := append([]string(nil), ids...)
	phOrig := append([]string(nil), phones...)
	_, _ = RedactAll(ids, phones)
	if ids[0] != idOrig[0] || phones[0] != phOrig[0] {
		t.Fatalf("RedactAll corrupted originals: %+v %+v", ids, phones)
	}
}
