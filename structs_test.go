package diff_test

import (
	"testing"

	"github.com/u-s-p/diff"
)

func TestStructs(t *testing.T) {
	s := diff.Elements{
		Left: []diff.Element{
			diff.Element{
				Id:      "hello",
				Content: []string{},
			},
			diff.Element{
				Id:      "world",
				Content: []string{},
			},
			diff.Element{
				Id:      "content",
				Content: []string{"hello", "world"},
			},
		},
		Right: []diff.Element{
			diff.Element{
				Id:      "its",
				Content: []string{},
			},
			diff.Element{
				Id:      "my",
				Content: []string{},
			},
			diff.Element{
				Id:      "world",
				Content: []string{},
			},
			diff.Element{
				Id:      "content",
				Content: []string{"its", "my", "world"},
			},
		},
	}
	l, r := s.Length()
	if l != 3 {
		t.Fatalf("Wrong left length, expected 3, got %d", l)
	}
	if r != 4 {
		t.Fatalf("Wrong right length, expected 4, got %d", r)
	}

	if s.Equal(0, 0) == diff.True {
		t.Fatalf("Expect to not equal")
	}
	if s.Equal(1, 1) == diff.True {
		t.Fatalf("Expect to not equal")
	}
	if s.Equal(1, 2) != diff.True {
		t.Fatalf("Expected equal")
	}
	if s.Equal(2, 3) != diff.Identity {
		t.Fatalf("Expected identity")
	}
}

func TestStructs2(t *testing.T) {
	s := diff.Elements{
		Left: []diff.Element{
			diff.Element{
				Id:      "content",
				Content: []string{"hello", "world"},
			},
			diff.Element{
				Id:      "hello",
				Content: []string{},
			},
		},
		Right: []diff.Element{
			diff.Element{
				Id:      "content",
				Content: []string{"its", "my", "world"},
			},
		},
	}
	l, r := s.Length()
	if l != 2 {
		t.Fatalf("Wrong left length, expected 2, got %d", l)
	}
	if r != 1 {
		t.Fatalf("Wrong right length, expected 1, got %d", r)
	}

	if s.Equal(0, 0) != diff.Identity {
		t.Fatalf("Expect identity")
	}
	if s.Equal(1, 0) != diff.False {
		t.Fatalf("Expect to not equal")
	}
	d := diff.New(s)
	if d[0].Delta != diff.Content {
		t.Fatalf("Expect to be ContentDiff")
	}
	if d[1].Delta != diff.Left {
		t.Fatalf("Expect to be Left")
	}
}