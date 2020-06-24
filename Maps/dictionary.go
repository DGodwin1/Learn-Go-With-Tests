package Maps

// 'Thin wrapper' on map that has string keys and values
type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	value, ok := d[word]
	if !ok {
		return "word does not exist"
	}
	return value
}
