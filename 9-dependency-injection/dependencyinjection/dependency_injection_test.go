package dependencyinjection_test

import (
	"bytes"
	"testing"

	. "github.com/louischering/learn-go-with-tests/9-dependency-injection/dependencyinjection"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
