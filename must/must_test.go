package must

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.NotPanics(t, func() {
		fn := func() error { return nil }
		Must(fn())
	})

	assert.Panics(t, func() {
		fn := func() error { return os.ErrNotExist }
		Must(fn())
	})
}

func TestMust2(t *testing.T) {
	assert.NotPanics(t, func() {
		fn := func() (string, error) { return "test", nil }
		assert.Equal(t, "test", Must2(fn()))
	})

	assert.Panics(t, func() {
		fn := func() (string, error) { return "test", os.ErrNotExist }
		Must2(fn())
	})
}
