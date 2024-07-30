package maps_and_tests

import "testing"

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

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("add a new word", func(t *testing.T) {
		dictionary.Add("test2", "another test")

		got, error := dictionary.Search("test2")
		want := "another test"

		if error != nil {
			t.Fatal("didn't expect to get an error")
		}

		assertStrings(t, got, want)
	})

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
