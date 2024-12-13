package pflagx

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert.Implements(t, (*flag.Value)(nil), &URL{})
}
