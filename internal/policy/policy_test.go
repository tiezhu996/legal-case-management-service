package policy

import "testing"

func containsString(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

func TestNewProviderNil(t *testing.T) {
	p := NewProvider(nil)
	if p != nil {
		t.Fatal("expected nil provider for nil rules")
	}
	if CheckAccess(p, "admin", "case", "read") {
		t.Fatal("nil provider should deny all")
	}
}

func TestCheckAccessDeniesNilProvider(t *testing.T) {
	var p PolicyProvider
	if CheckAccess(p, "admin", "case", "read") {
		t.Fatal("nil provider should deny")
	}
}

func TestPolicyLoaderAdd(t *testing.T) {
	l := NewPolicyLoader(nil)
	l.Add("lawyer", "case", []string{"read", "update"})
	if !containsString(l.Actions("lawyer", "case"), "read") {
		t.Fatal("missing read action after Add")
	}
}

func TestPolicyLoaderActionsUnknown(t *testing.T) {
	l := NewPolicyLoader(nil)
	if l.Actions("lawyer", "case") != nil {
		t.Fatal("expected nil actions for unknown entry")
	}
}
