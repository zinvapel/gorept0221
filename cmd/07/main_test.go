package main

import "testing"

func TestDo(t *testing.T) {
	if err := Do(); err != nil {
		t.Error("unexpected error", err)
	}
}