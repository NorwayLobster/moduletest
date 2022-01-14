package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHello1(t *testing.T) {
	want := "Hi, World"
	if got := Hello1(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	want := "Hello World"
	if got := Hello2(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "proverb"
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
