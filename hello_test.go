package moduletest

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
