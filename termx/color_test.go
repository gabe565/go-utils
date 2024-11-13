package termx

import (
	"testing"

	"github.com/creack/pty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsColor(t *testing.T) {
	pty, tty, err := pty.Open()
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = pty.Close()
		_ = tty.Close()
	})

	t.Run("true", func(t *testing.T) {
		t.Setenv("TERM", "xterm")
		t.Setenv("NO_COLOR", "")
		assert.True(t, IsColor(pty))
	})

	t.Run("false when TERM=dumb", func(t *testing.T) {
		t.Setenv("TERM", "dumb")
		t.Setenv("NO_COLOR", "")
		assert.False(t, IsColor(pty))
	})

	t.Run("false when NO_COLOR=true", func(t *testing.T) {
		t.Setenv("TERM", "xterm")
		t.Setenv("NO_COLOR", "true")
		assert.False(t, IsColor(pty))
	})
}
