package main

type Dictionary map[string]string

const (
	ErrWordExists        = DictionaryErr("the word you were looking for already exists")
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrNotFoundForUpdate = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// maps can return two values including a found/not found boolean
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFoundForUpdate
	case nil:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
