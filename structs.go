package diff

import (
	"fmt"
	"strings"
)

type Element struct {
	Id      string
	Content []string
}

// Holds two Element lists to diff.
type Elements struct {
	Left, Right []Element
}

func (e Elements) Equal(left, right int) Equal {
	switch {
	case e.Left[left].Id == e.Right[right].Id && strings.Join(e.Left[left].Content, ",") == strings.Join(e.Right[right].Content, ","):
		return True
	case e.Left[left].Id == e.Right[right].Id:
		return Identity
	default:
		return False
	}
}

func (e Elements) Length() (int, int) {
	return len(e.Left), len(e.Right)
}

func (e Elements) ContentDiff(left, right int) (iface Interface) {
	return Strings{
		Left:  e.Left[left].Content,
		Right: e.Right[right].Content,
	}
}

func (e Element) String() string {
	return fmt.Sprintf("ID: %s\n%s", e.Id, e.Content)
}

func prefixBlock(pfx, str string) string {
	strs := strings.Split(str, "\n")
	for i, s := range strs {
		strs[i] = pfx + s
	}
	return strings.Join(strs, "\n")
}
