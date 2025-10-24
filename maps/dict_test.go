package dict

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("search found", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		dict := Dictionary{
			key: value,
		}

		assertDictValue(t, dict, key, value)
	})

	t.Run("search not found", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		dict := Dictionary{
			key: value,
		}

		got, err := dict.Search("invalid")
		expected := ""
		assertError(t, err, ErrKeyNotFound)
		assertStrings(t, got, expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new key", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		dict := Dictionary{}

		err := dict.Add(key, value)

		assertNotError(t, err)
		assertDictValue(t, dict, key, value)
	})

	t.Run("add old key", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		newValue := "This is new"
		dict := Dictionary{
			key: value,
		}

		err := dict.Add(key, newValue)

		assertError(t, err, ErrDuplicateKey)
		assertDictValue(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update old key", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		dict := Dictionary{
			key: value,
		}
		newValue := "This is new value"

		err := dict.Update(key, newValue)

		assertNotError(t, err)
		assertDictValue(t, dict, key, newValue)
	})

	t.Run("update new key", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "This is a test"

		err := dict.Update(key, value)
		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete old key", func(t *testing.T) {
		key := "test"
		value := "This is a test"
		dict := Dictionary{
			key: value,
		}

		err := dict.Delete(key)
		assertNotError(t, err)

		got, err := dict.Search(key)
		assertError(t, err, ErrKeyNotFound)
		assertStrings(t, got, "")
	})

	t.Run("delete new key", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"

		err := dict.Delete(key)

		assertError(t, err, ErrKeyNotFound)
	})
}

func assertStrings(tb testing.TB, got, expected string) {
	tb.Helper()
	if got != expected {
		tb.Errorf("got %q but expected %q", got, expected)
	}
}

func assertDictValue(tb testing.TB, dict Dictionary, key, value string) {
	tb.Helper()

	got, err := dict.Search(key)
	expected := value

	assertNotError(tb, err)
	assertStrings(tb, got, expected)
}

func assertNotError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Errorf("got %q but expected none", err)
	}
}

func assertError(tb testing.TB, got, expected error) {
	tb.Helper()
	if !errors.Is(got, expected) {
		tb.Errorf("got %q but expected %q", got, expected)
	}
}
