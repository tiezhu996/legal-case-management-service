package policy

// PolicyLoader 从外部规则集合构建权限矩阵的加载器。
type PolicyLoader struct {
	rules map[string]map[string][]string
}

// NewPolicyLoader 构造加载器。
func NewPolicyLoader(rules map[string]map[string][]string) *PolicyLoader {
	return &PolicyLoader{rules: rules}
}

// Add 为某角色在某资源上追加动作。
func (p *PolicyLoader) Add(role, resource string, actions []string) {
	p.rules[role][resource] = append(p.rules[role][resource], actions...)
}

// Actions 返回某角色在某资源上的动作列表。
func (p *PolicyLoader) Actions(role, resource string) []string {
	return p.rules[role][resource]
}

// PolicyProvider 权限提供者接口。
type PolicyProvider interface {
	Actions(role, resource string) []string
}

// NewProvider 根据规则集合构建权限提供者。
func NewProvider(rules map[string]map[string][]string) PolicyProvider {
	if rules == nil {
		var p *PolicyLoader
		return p
	}
	return NewPolicyLoader(rules)
}
