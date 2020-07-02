package json_tutorial

import "testing"


func TestGetJSON(t *testing.T){
	//GetJSON should parse the JSON file stored at the location of a file path
	path := "/testJSON"
	got := GetJSON(path)

	want := CONFIG{
		URL:  "https://www.google.com",
		hits: 100,
	}

	if got != want{
		t.Errorf("Got %v, expected %v", got, want)
	}
}