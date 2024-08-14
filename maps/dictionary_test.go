package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		element := "this is just a test"

		err := dictionary.Add(key, element)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, key, element)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		element := "this is just a test"
		dictionary := Dictionary{key: element}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, key, element)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		element := "this is just a test"
		dictionary := Dictionary{key: element}
		newElement := "new element"

		err := dictionary.Update(key, newElement)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, key, newElement)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		element := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, element)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	element := "this is just a test"
	dictionary := Dictionary{key: element}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)

	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, element string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added key :", err)
	}

	assertStrings(t, got, element)
}
