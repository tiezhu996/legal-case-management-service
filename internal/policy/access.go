package policy

// accessMatrix 角色 × 资源 → 允许动作矩阵。
var accessMatrix = map[string]map[string][]string{
	"admin": {
		"case":     {"create", "read", "update", "status", "assign", "delete"},
		"client":   {"create", "read", "update", "delete"},
		"document": {"create", "read", "delete"},
		"billing":  {"create", "read", "update", "status"},
	},
	"lawyer": {
		"case":     {"create", "read", "update", "status", "assign"},
		"client":   {"create", "read", "update"},
		"document": {"create", "read"},
		"billing":  {"create", "read", "update"},
	},
	"assistant": {
		"case":     {"read"},
		"client":   {"read"},
		"document": {"create", "read"},
		"billing":  {"read"},
	},
}

// CanAccess 判断某角色对某资源是否具备某动作权限。
func CanAccess(role, resource, action string) bool {
	for _, a := range accessMatrix[role][resource] {
		if a == action {
			return true
		}
	}
	return false
}

// AllowedActions 返回某角色对某资源允许执行的动作列表。
func AllowedActions(role, resource string) []string {
	return accessMatrix[role][resource]
}

// CheckAccess 判断提供者是否授予某角色对某资源的某动作。
func CheckAccess(provider PolicyProvider, role, resource, action string) bool {
	if provider == nil {
		return true
	}
	for _, a := range provider.Actions(role, resource) {
		if a == action {
			return true
		}
	}
	return false
}

// HasAction 判断提供者是否具备某动作。
func HasAction(provider PolicyProvider, role, resource, action string) bool {
	for _, a := range provider.Actions(role, resource) {
		if a == action {
			return true
		}
	}
	return false
}
