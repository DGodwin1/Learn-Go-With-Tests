package Maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "test"
		got := dictionary.Search(key)
		want := "this is just a test"

		SearchHelper(t, got, want, key)
	})

	t.Run("word not in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"word": "a value"}
		key := "hello"
		got := dictionary.Search(key)
		want := "word does not exist"
		SearchHelper(t, got, want, key)
	})
}

func SearchHelper(t *testing.T, got string, want string, key string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q but wanted %q given key: %q", got, want, key)
	}

}
