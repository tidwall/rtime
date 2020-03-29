package rtime

import (
	"testing"
)

func TestTime(t *testing.T) {
	tm := Now()
	if tm.IsZero() {
		t.Fatal("zero time")
	}
}
