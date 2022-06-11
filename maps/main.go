package main

const (
	ErrWordNotExist DictionaryErr = "the target word does not exist in dictionary"
	ErrAlreadyAdded DictionaryErr = "cannot add word because it is already added in dictionary"
)

type Dictionary map[string]string

type DictionaryErr string

func (d Dictionary) Search(word string) (string, error) {
	if _, ok := d[word]; !ok {
		return "", ErrWordNotExist
	}
	return d[word], nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrAlreadyAdded
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; !ok {
		return ErrWordNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
