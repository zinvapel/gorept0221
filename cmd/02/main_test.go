package main

import (
	"bytes"
	"github.com/zinvapel/gorept0221/pkg/two"
	"log"
	"testing"
)

var buf = &bytes.Buffer{}

func TestMain(m *testing.M) {
	log.SetOutput(buf)
	log.SetFlags(0)
	m.Run()
}

func TestDo(t *testing.T) {
	Do()

	switch two.GetFilePermission() {
	case 0:
		assertResult(t, "")
	case 1:
		assertResult(t, "Can execute\n")
	case 2:
		assertResult(t, "Can write\n")
	case 3:
		assertResult(t, "Can read\nCan execute\n")
	case 4:
		assertResult(t, "Can read\n")
	case 5:
		assertResult(t, "Can write\nCan execute\n")
	case 6:
		assertResult(t, "Can read\nCan write\n")
	case 7:
		assertResult(t, "Can read\nCan write\nCan execute\n")
	}
}

func assertResult(t *testing.T, expected string) {
	if buf.String() != expected {
		t.Errorf("Unexpected result: '%s', expected: '%s'", buf.String(), expected)
	}
}