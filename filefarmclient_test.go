package filefarmclient

import (
	"testing"
)

func TestHello(t *testing.T) {
	got, err := Hello()
	if err != nil {
		t.Errorf("Hello() encounters error %q", err.Error())
	}
	if got != "hello" {
		t.Errorf("Hello() == %q, want %q", got, "hello")
	}
}
