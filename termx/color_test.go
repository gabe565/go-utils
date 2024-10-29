package termx

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/creack/pty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsTerminal(t *testing.T) {
	pty, tty, err := pty.Open()
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = pty.Close()
		_ = tty.Close()
	})

	t.Run("true for PTY", func(t *testing.T) {
		assert.True(t, IsTerminal(pty))
	})

	t.Run("false for io.Discard", func(t *testing.T) {
		assert.False(t, IsTerminal(io.Discard))
	})

	t.Run("false for *os.File", func(t *testing.T) {
		tmp := t.TempDir()
		f, err := os.Create(filepath.Join(tmp, "test"))
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = f.Close()
		})
		assert.False(t, IsTerminal(f))
	})
}

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
