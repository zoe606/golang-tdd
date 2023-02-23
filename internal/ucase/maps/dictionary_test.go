package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	//dictionary := map[string]int{"test": 10}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"
		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		def := "this is just a test"
		word := "test"

		err := dic.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dic := Dictionary{word: def}

		err := dic.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dic := Dictionary{word: def}
		newDef := "new definition"

		err := dic.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "word"
		def := "this is just a test"
		dic := Dictionary{}

		err := dic.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dic Dictionary, word string, def string) {
	t.Helper()
	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if def != got {
		t.Errorf("got %q want %q", got, def)
	}
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionary{word: "test of definition"}

	dic.Delete(word)
	_, err := dic.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got error %q expected %q", got, expected)
	}
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q give, %q", got, expected, "test")
	}
}
