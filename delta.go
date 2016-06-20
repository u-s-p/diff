package diff

type Equal int

const (
	True Equal = iota
	Identity
	False
)

func (equal Equal) String() string {
	switch equal {
	case True:
		return "true"
	case Identity:
		return "identity"
	case False:
		return "false"
	}
	return "unknown"
}

// Describe in which collection the element occurs; Left, Right or Both.
type Delta int

const (
	// Element is present in both Left and Right collections.
	// Index uses the Left collection.
	Both Delta = iota
	// Element is present only in the Left collection.
	// Index uses the Left collection.
	Left
	// Element is present only in the Right collection.
	// Index uses the Right collection.
	Right
	// Content is different
	Content
)

func (delta Delta) String() string {
	switch delta {
	case Both:
		return "Both"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Content:
		return "ContentDiff"
	}
	return "unknown"
}

// One Diff record per element.
// If Delta is Left or Both, Index is for the left collection.
// If Delta is Right, Index is for the right collection.
type Diff struct {
	Delta       Delta
	Index       int
	IndexR      int
	ContentDiff []Diff
}
