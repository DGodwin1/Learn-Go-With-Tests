package json_tutorial

import "testing"

func TestJsonEncodesToStruct(t *testing.T){

	got := ParseJSON()

	want := &Config{"https://www.google.com", 100}

	if got.URL != want.URL{
		t.Errorf("Got %v, want %v", got.URL, want.URL)
	}

	if got.Hits != want.Hits{
		t.Errorf("Got %v, want %v", got.Hits, want.Hits)
	}

}
