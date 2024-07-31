package maps_and_tests

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("couldn't find the word you're looking for")
	ErrWordExists       = DictionaryErr("word is already defined")
	ErrWordDoesNotExist = DictionaryErr("couldn't update word, as the word doesn't exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(target string) (string, error) {

	// map lookup can return 2 values, the 2nd is a boolean which indicates if the key was found successfully
	// This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.
	definition, ok := d[target]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
