package maps

import "errors"

type Dictionary map[string]string

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
