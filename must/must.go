package must

// Must is a helper that wraps a call to a function returning (error)
// and panics if the error is non-nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must2 is a helper that wraps a call to a function returning (T, error)
// and panics if the error is non-nil.
func Must2[T any](val T, err error) T { //nolint:ireturn,nolintlint
	Must(err)
	return val
}
