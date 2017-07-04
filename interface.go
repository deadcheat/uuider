package uuider

// Producer UUID Producer interface
type Producer interface {
	New() string
}
