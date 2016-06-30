package diff

type NoContent struct {
	Interface
}

func (nc NoContent) Equal(left, right int) bool {
	return nc.Interface.Equal(left, right)
}

func (nc NoContent) EqualContent(left, right int) Equal {
	if nc.Interface.Equal(left, right) {
		return True
	}
	return False
}

func (nc NoContent) Length() (left, right int) {
	return nc.Interface.Length()
}

func (nc NoContent) ContentDiff(left, right int) (iface InterfaceContent) {
	return NoContent{}
}

// Runs a diff on the given Interface.
// Returns the results as a slice of Diff.
func New(iface Interface) []Diff {
	return NewContent(NoContent{iface})
}

// Runs a diff on the given Interface.
// Returns the results as a slice of Diff.
func NewContent(iface InterfaceContent) []Diff {
	l, r := iface.Length()
	diff := make([]Diff, 0, l+r)
	diff = snipEnd(iface, l, r, diff)
	l -= len(diff)
	r -= len(diff)
	if l != 0 && r != 0 {
		table := lcs(iface, l, r)
		diff = walk(iface, l, r, table, diff)
	} else if l != 0 {
		diff = remainingOneSide(iface, l, Left, diff)
	} else if r != 0 {
		diff = remainingOneSide(iface, r, Right, diff)
	}
	reverse(diff)
	return diff
}

// Handle the identical Left and Right tails.
func snipEnd(iface InterfaceContent, l, r int, diff []Diff) []Diff {
	min := l
	if r < min {
		min = r
	}
loop:
	for i := 0; i < min; i++ {
		switch iface.EqualContent(l-1-i, r-1-i) {
		case True:
			diff = append(diff, Diff{Delta: Both, Index: l - 1 - i})
		default:
			break loop
		}
	}
	return diff
}

// Handle the rest of the diff, if one of the two collections is exhausted after snipEnd.
func remainingOneSide(iface InterfaceContent, idx int, delta Delta, diff []Diff) []Diff {
	for i := 0; i < idx; i++ {
		diff = append(diff, Diff{Delta: delta, Index: idx - 1 - i})
	}
	return diff
}

// Constructs a LCSLength table
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Computing_the_length_of_the_LCS
func lcs(iface InterfaceContent, l, r int) [][]int {
	lnum, rnum := l, r
	rows, cols := lnum+1, rnum+1
	table := make([][]int, rows)
	cels := make([]int, rows*cols)
	for i := 0; i < rows; i++ {
		table[i] = cels[:cols]
		cels = cels[cols:]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			switch iface.EqualContent(i-1, j-1) {
			case True, Identity:
				table[i][j] = table[i-1][j-1] + 1
			default:
				a := table[i-1][j]
				b := table[i][j-1]
				if b > a {
					a = b
				}
				table[i][j] = a
			}
		}
	}
	return table
}

// Walk the lcs table
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Example
func walk(iface InterfaceContent, l, r int, table [][]int, diff []Diff) []Diff {
	i, j := l, r
	for {
		if i == 0 && j == 0 {
			return diff
		} else if i == 0 {
			j--
			diff = append(diff, Diff{Delta: Right, Index: j})
		} else if j == 0 {
			i--
			diff = append(diff, Diff{Delta: Left, Index: i})
		} else {
			switch iface.EqualContent(i-1, j-1) {
			case True:
				i--
				j--
				diff = append(diff, Diff{Delta: Both, Index: i})
			case Identity:
				i--
				j--
				diff = append(diff, Diff{Delta: Content, Index: i, IndexR: j, ContentDiff: NewContent(iface.ContentDiff(i, j))})
			default:
				if table[i-1][j] > table[i][j-1] {
					i--
					diff = append(diff, Diff{Delta: Left, Index: i})
				} else {
					j--
					diff = append(diff, Diff{Delta: Right, Index: j})
				}
			}
		}
	}
}

func reverse(diff []Diff) {
	i := 0
	j := len(diff) - 1
	for i < j {
		diff[i], diff[j] = diff[j], diff[i]
		i++
		j--
	}
}
