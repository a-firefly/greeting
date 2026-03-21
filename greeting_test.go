package greeting

import "testing"

func TestHello(t *testing.T) {
	got := Greet("Bob", "Nice to meet you!")
	want := "Hello, Bob. Nice to meet you!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
