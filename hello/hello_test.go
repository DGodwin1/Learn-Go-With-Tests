package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David"

		assertCorrect(t, got, want)

	})

	t.Run("say 'Hello world' if empty name string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrect(t, got, want)

	})

	t.Run("Give back Spanish reply if Spanish specified", func(t *testing.T) {
		got := Hello("David", "Spanish")
		want := "Hola, David"

		assertCorrect(t, got, want)

	})

	t.Run("Give back Hola World if Spanish specified and empty string for name", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		assertCorrect(t, got, want)

	})

	t.Run("Give back Bonjour if French specified and empty string for name", func(t *testing.T){
		got := Hello("","French")
		want := "Bonjour, World"

		assertCorrect(t, got, want)

	})

	t.Run("Give back Bonjour + name if French specified", func(t *testing.T){
		got := Hello("David","French")
		want := "Bonjour, David"

		assertCorrect(t, got, want)

	})

	t.Run("Give back 'Leave me alone' if language is Rude", func(t *testing.T){
		got := Hello("David","Rude")
		want := "Leave me alone, David"

		assertCorrect(t, got, want)

	})

	t.Run("Give back 'Leave me alone, World' if language is Rude but no name", func(t *testing.T){
		got := Hello("", "Rude")
		want := "Leave me alone, World"
		assertCorrect(t, got, want)
	})

}
