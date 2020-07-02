package json_tutorial

import (
	"reflect"
	"testing"
)

func TestJsonEncodesToStruct(t *testing.T){
	got := ParseJSON()

	want := &Config{"https://www.google.com", 100}

	if !reflect.DeepEqual(got, want){
		t.Errorf("Got %v, want %v", got, want)
	}
}
