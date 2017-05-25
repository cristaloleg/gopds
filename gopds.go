package gopds

// Container basic interface
// for all persistent data structures
type Container interface {
	Size() int
	IsEmpty() bool
}
