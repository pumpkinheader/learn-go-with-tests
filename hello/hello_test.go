package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say 'HEllo, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("mikajui", "")
		want := "Hello, mikajui"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got:= Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got:= Hello("Hiroyuki", "French")
		want := "Bonjour, Hiroyuki"
		assertCorrectMessage(t, got, want)
	})
}