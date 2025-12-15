package weather

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"
	if got != want {
		t.Errorf(`Hello() = %q, want "%q", error`, got, want)
	}
}
