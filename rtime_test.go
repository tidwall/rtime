package rtime

import (
	"testing"
	"time"
)

func unsync() {
	smu.Lock()
	synced = false
	smu.Unlock()
}

func TestTime(t *testing.T) {
	unsync()
	start := time.Now()
	tm := Now()

	if tm.IsZero() {
		t.Fatal("zero time")
	}
	println(time.Since(start).String())
}

func TestSync(t *testing.T) {
	unsync()
	if err := Sync(); err != nil {
		t.Fatal(err)
	}
	if !synced {
		t.Fatal("not synced")
	}
	tm1 := Now()
	if tm1.IsZero() {
		t.Fatal("zero time")
	}
	tm2 := Now()
	if !tm2.After(tm1) {
		t.Fatal("not after")
	}
	func() {
		defer func() {
			if v := recover(); v != nil {
				t.Fatal(v)
			}
		}()
		MustSync()
	}()

}
