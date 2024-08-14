package dictionary

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(element string) (string, error) {
	definition, ok := d[element]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, element string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = element
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
