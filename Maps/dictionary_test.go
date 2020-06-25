package Maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "test"
		got, _ := dictionary.Search(key)
		want := "this is just a test"

		SearchHelper(t, got, want, key)
	})

	t.Run("word not in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"word": "a value"}
		key := "hello"
		_, got := dictionary.Search(key)
		SearchHelper(t, got.Error(), ErrWordNotFound.Error(), key)
	})
}

func TestAdd(t *testing.T){
	t.Run("Add word that doesn't exist", func(t *testing.T){
		dictionary := Dictionary{}
		word := "budweiser"
		definition := "beer"
		dictionary.Add(word,definition)

		// Now let's find the word in the dictionary
		assertDefinition(t, dictionary, word, definition)

	})
	t.Run("Add word that already exists", func(t *testing.T){
		dictionary := Dictionary{"budweiser":"beer"}

		// Refuse to add word that already exists.
		err := dictionary.Add("budweiser","drink")

		if err == nil{
			t.Fatal("should have got an error", err)
		}
		// Definition should be for beer.
		assertDefinition(t, dictionary, "budweiser", "beer")
	})
}

func TestUpdate(t *testing.T){
	t.Run("Update a word that already exists", func(t *testing.T) {
		dictionary := Dictionary{"budweiser":"drink"}
		err := dictionary.Update("budweiser","beer")
		if err != nil{
			t.Fatal("shouldn't have gotten an error", err)
		}
		assertDefinition(t, dictionary, "budweiser", "beer")

	})

	}
func TestDelete(t *testing.T) {
	t.Run("Delete a word from dictionary", func(t *testing.T) {
		word := "budweiser"
		dictionary := Dictionary{word: "drink"}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrWordNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}

	})
}

func SearchHelper(t *testing.T, got string, want string, key string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q but wanted %q given key: %q", got, want, key)
	}

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}