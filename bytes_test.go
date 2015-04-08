package diff_test

import (
	"testing"

	"github.com/PieterD/diff"
)

func TestBytes(t *testing.T) {
	s := diff.Bytes{
		Left:  [][]byte{[]byte("hello"), []byte("world")},
		Right: [][]byte{[]byte("its"), []byte("my"), []byte("world")},
	}
	l, r := s.Length()
	if l != 2 {
		t.Fatalf("Wrong left length, expected 2, got %d", l)
	}
	if r != 3 {
		t.Fatalf("Wrong right length, expected 3, got %d", r)
	}

	if s.Equal(0, 0) {
		t.Fatalf("Did not expect equal")
	}
	if s.Equal(1, 1) {
		t.Fatalf("Did not expect equal")
	}
	if !s.Equal(1, 2) {
		t.Fatalf("Expected equal")
	}
}
