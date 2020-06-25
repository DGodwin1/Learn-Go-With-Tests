package Maps

import "errors"

// 'Thin wrapper' on map that has string keys and values.
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
	_, err := d.Search(key)

	// If the word already exists in the dictionary
	// then you shouldn't overwrite the value that
	// is already there.
	if err == nil{
		return errors.New("word already exists")
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	if err != nil{
		errors.New("Word not found so can't update it")
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d,word)
}


