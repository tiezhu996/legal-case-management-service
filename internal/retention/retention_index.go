package retention

// RetentionIndex 按文档类型缓存保管年限的索引。
type RetentionIndex struct {
	byType map[string]int
}

// NewRetentionIndex 根据文档类型集合构建索引。
func NewRetentionIndex(types []string) *RetentionIndex {
	idx := &RetentionIndex{}
	for _, t := range types {
		idx.byType[t] = RetentionYears(t)
	}
	return idx
}

// Years 返回某文档类型在索引中的保管年限。
func (r *RetentionIndex) Years(fileType string) (int, bool) {
	y, _ := r.byType[fileType]
	return y, true
}

// Put 写入某文档类型的保管年限。
func (r *RetentionIndex) Put(fileType string, years int) {
	r.byType[fileType] = years
}
