// @Author: 2014BDuck
// @Date: 2021/2/14

package dictionary

import "testing"

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, defination string) {
	t.Helper()

	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if defination != got {
		t.Errorf("got '%s' want '%s'", got, defination)
	}
}

func TestSearch(t *testing.T) {
	assertString := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
		}
	}

	dictionary := Dictionary{
		"test": "this is just a test",
	}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "this is just a test"
		err := dictionary.Add(word, defination)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrorWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("word not exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"

		newWord := "new test"
		err := dictionary.Update(newWord, newDefinition)
		assertError(t, err, ErrorNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	t.Run("delete exist", func(t *testing.T) {
		err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrorNotFound)
	})

	t.Run("delete not exist", func(t *testing.T) {
		err := dictionary.Delete(word)
		assertError(t, err, ErrorNotFound)
	})
}
