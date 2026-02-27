package greeting

import "testing"

func TestHello(t *testing.T) {
	got := Greet("Bob")
	want := "Hello, Bob!"
	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
