package main

import "testing"

func TestHello(t *testing.T){
	assertCorrect := func(t *testing.T, got, want string){
		t.Helper()
		if got != want{
			t.Errorf("got %q want %q",got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("David")
		want := "Hello, David"

		if got != want{
			assertCorrect(t, got, want)
		}
	})

	t.Run("say 'Hello world' if empty string supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, World"

		if got != want{
			assertCorrect(t, got, want)
		}
	})
}