package maps

import "errors"

const (
	NOT_IN_DICTIONARY_MESSAGE        = "could not find entry in dictionary"
	DUPLICATED_KEY_MESSAGE           = "word already exists within dictionary"
	WORD_DOESNT_EXIST_UPDATE_MESSAGE = "word does not exist to be updated"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", errors.New(NOT_IN_DICTIONARY_MESSAGE)
	}
	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return errors.New(DUPLICATED_KEY_MESSAGE)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return errors.New(WORD_DOESNT_EXIST_UPDATE_MESSAGE)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errors.New(NOT_IN_DICTIONARY_MESSAGE)
	}
	delete(d, word)
	return nil
}
