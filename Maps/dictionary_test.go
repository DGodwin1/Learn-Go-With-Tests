package Maps

import "testing"

func TestSearch(t *testing.T){
	// map[string]string == a map will contain a string key which
	// will deliver a string value.
	dictionary := map[string]string{"test":"this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want{
		t.Errorf("Got %q but wanted %q given key: %q", got, want, "test")
	}
}
