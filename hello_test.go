package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomName() string {
	names := []string{
		"Matt",
		"Elodie",
		"Marc",
		"Freddie",
	}
	randomIndex := rand.Intn(len(names))
	return names[randomIndex]
}

// Tests must begin with Test
func TestHello(t *testing.T) {
	assetCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // Tells test quite this func is a helper so errors will show the real line number and now this funcs
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		name := randomName()
		got := Hello(name, "")
		want := "Hello, " + name
		assetCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assetCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		name := randomName()
		got := Hello(name, "spanish")
		want := "Hola, " + name
		assetCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		name := randomName()
		got := Hello(name, "french")
		want := "Bonjour, " + name
		assetCorrectMessage(t, got, want)
	})
	t.Run("in mixed case language ", func(t *testing.T) {
		name := randomName()
		got := Hello(name, "SpaNIsh")
		want := "Hola, " + name
		assetCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	res := Hello("Sabastian", "")
	fmt.Println(res)
	// Output: Hello, Sabastian
}
