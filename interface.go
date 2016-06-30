package diff

// Wrap your data in one of these to diff it.
// It should hold two collections, the Left (or old) one and the Right (or new) one.
type Interface interface {
	// Return true if the elements at the given indices in Left and Right are equal.
	Equal(left, right int) (isEqual bool)
	// Return the sizes of the Left and Right collections.
	Length() (left int, right int)
}

type InterfaceContent interface {
	// Return true if the elements at the given indices in Left and Right are equal.
	Equal(left, right int) (isEqual bool)
	// Return if the elements at the given indices in Left and Right are equal, if the identity is equal or if the are different.
	EqualContent(left, right int) (isEqual Equal)
	// Return the sizes of the Left and Right collections.
	Length() (left, right int)
	// Process the diff operation on the content.
	ContentDiff(left, right int) (iface InterfaceContent)
}
