package dict

import (
	"errors"
)

const (
	ErrKeyNotFound  = ErrDict("key not found")
	ErrDuplicateKey = ErrDict("duplicate key")
)

type Dictionary map[string]string

type ErrDict string

func (e ErrDict) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch {
	case err == nil:
		return ErrDuplicateKey
	case errors.Is(err, ErrKeyNotFound):
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch {
	case err == nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch {
	case err == nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
