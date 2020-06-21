package Iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	got := Repeat("a", 4)
	want := "aaaa"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat(){
	result := Repeat("a", 4)
	fmt.Println(result)
	// Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a",4)
	}
}