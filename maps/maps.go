package maps_and_tests

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("couldn't find the word you're looking for")

func (d Dictionary) Search(target string) (string, error) {

	// map lookup can return 2 values, the 2nd is a boolean which indicates if the key was found successfully
	// This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.
	definition, ok := d[target]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
