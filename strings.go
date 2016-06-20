package diff

// Holds two string lists to diff.
type Strings struct {
	Left, Right []string
}

func (str Strings) Equal(left, right int) Equal {
	if str.Left[left] == str.Right[right] {
		return True
	}
	return False
}

func (str Strings) Length() (int, int) {
	return len(str.Left), len(str.Right)
}

func (str Strings) ContentDiff(left, right int) (iface Interface) {
	return Strings{}
}