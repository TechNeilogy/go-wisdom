package typew

// ConvertThenAssert converts t to any, then type asserts it to U.
func ConvertThenAssert[U any, T any](t T) (U, bool) {
	a, ok := any(t).(U)
	return a, ok
}
