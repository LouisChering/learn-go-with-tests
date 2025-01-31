package maps_test

import (
	"testing"

	. "github.com/louischering/learn-go-with-tests/7-maps"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := NOT_IN_DICTIONARY_MESSAGE

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}
func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		assertError(t, err)
		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})
	t.Run("should not overwrite existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is a test")
		assertError(t, err)
		err = dictionary.Add("test", "this is a test")
		if err == nil {
			t.Error("expected an error")
			t.FailNow()
		}
		if err.Error() != DUPLICATED_KEY_MESSAGE {
			t.Errorf("expected %s but got %s", DUPLICATED_KEY_MESSAGE, err.Error())
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		if err != nil && err.Error() != WORD_DOESNT_EXIST_UPDATE_MESSAGE {
			t.Errorf("expected %q but got %q", WORD_DOESNT_EXIST_UPDATE_MESSAGE, err.Error())
		}
	})

}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)

		assertError(t, err)

		_, err = dictionary.Search(word)

		if err != nil && err.Error() != NOT_IN_DICTIONARY_MESSAGE {
			t.Errorf("expected %q but got %q", NOT_IN_DICTIONARY_MESSAGE, err.Error())
		}
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		if err != nil && err.Error() != NOT_IN_DICTIONARY_MESSAGE {
			t.Errorf("expected %q but got %q", NOT_IN_DICTIONARY_MESSAGE, err.Error())
		}
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error but got %s", err.Error())
	}
}
