package store

// A secret without any metadata
type RawSecret struct {
	Value string
	Key   string
}
