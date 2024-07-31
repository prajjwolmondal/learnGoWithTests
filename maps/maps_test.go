package maps_and_tests

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search for existing key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("search for non-existent key", func(t *testing.T) {
		_, err := dictionary.Search("err")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAddingWords(t *testing.T) {

	t.Run("add a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "another test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add an existing word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "a new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdatingWords(t *testing.T) {

	t.Run("update definition of existing word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{word: definition}
		updatedDefinition := "newly updated definition"

		dictionary.Update(word, updatedDefinition)
		assertDefinition(t, dictionary, word, updatedDefinition)
	})

	t.Run("update definition of new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "definition")

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDeleteWords(t *testing.T) {

	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, error := dictionary.Search(word)

	if error != nil {
		t.Fatal("didn't expect to get an error")
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
