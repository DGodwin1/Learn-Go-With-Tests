package Integers

import "testing"

func TestAdder(t *testing.T){
	sum := Add(2,2)
	want := 4

	if want != sum{
		t.Errorf("expected '%d' but got '%d'", want, sum)
	}
}