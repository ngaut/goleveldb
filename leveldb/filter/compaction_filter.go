package filter

// CompactionFilter is used for compaction
type CompactionFilter interface {
	// The name of the compaction filter, for logging
	Name() string

	// The kv should be preserved If the Filter method returns false,
	// otherwise the kv should be removed
	// To modify the existing value, return true and newVal.
	// To retain the previous value, simply return nil
	//
	// The application must ensure that the call is thread-safe.
	Filter(key, val []byte) (remove bool, newVal []byte)
}
