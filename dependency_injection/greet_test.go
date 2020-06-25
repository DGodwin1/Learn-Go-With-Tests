package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "David")

	got := buffer.String()
	want := "Hello, David"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}
