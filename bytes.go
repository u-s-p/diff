package diff

import "bytes"

// Holds two string lists to diff.
type Bytes struct {
	Left, Right [][]byte
}

func (str Bytes) Equal(left, right int) Equal {
	if bytes.Equal(str.Left[left], str.Right[right]) {
		return True
	}
	return False
}

func (str Bytes) Length() (int, int) {
	return len(str.Left), len(str.Right)
}

func (str Bytes) ContentDiff(left, right int) (iface Interface) {
	return Bytes{}
}
