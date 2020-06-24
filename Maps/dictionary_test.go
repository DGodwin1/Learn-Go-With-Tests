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
	t.Run("Add word to dictionary", func(t *testing.T){
		dictionary := Dictionary{}
		dictionary.Add("budweiser","beer")

		// Now let's find the word in the dictionary
		got, err := dictionary.Search("budweiser")
		want := "beer"

		if err != nil{
			t.Fatal("should find added word:", err)
		}

		SearchHelper(t, got, want, "budweiser")

	})
}

func SearchHelper(t *testing.T, got string, want string, key string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q but wanted %q given key: %q", got, want, key)
	}

}
