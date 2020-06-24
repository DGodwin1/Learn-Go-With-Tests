package Maps

import "errors"

// 'Thin wrapper' on map that has string keys and values
type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not found")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error{
	d[key] = value
	return nil
}


