package utils

import (
	"testing"
)

func AssertEqual[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected == actual {
		return
	}

}
