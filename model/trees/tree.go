package redblacktree

type Tree interface {
	Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}

// Container is base interface that all data structures implement.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
