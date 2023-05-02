package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("mikajui")
	want := "Hello, mikajui"

	if	got != want {
		t.Errorf("got %q want %q", got , want)
	}
}