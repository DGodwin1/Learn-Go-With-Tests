package Maps

func Search(m map[string]string, k string) string{
	value, ok := m[k]
	if ok{
		return value
	}
	return ""
}
