package sagam

import (
	"bytes"
	"testing"
	"time"

	"github.com/obzva/timer"
)

func TestOpening(t *testing.T) {
	var buffer bytes.Buffer

	printOpening(&buffer)

	got := buffer.String()
	want := openingText

	assertString(t, got, want)
}

func TestSprint(t *testing.T) {
	timer := timer.NewTimer(time.Nanosecond)
	timer.Start()

	got := sprint(timer)
	assertString(t, got, statusRest)
}

func TestRest(t *testing.T) {
	timer := timer.NewTimer(time.Nanosecond)
	timer.Start()
	got := rest(timer)
	assertString(t, got, statusSprint)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
