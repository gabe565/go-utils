package termx

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/creack/pty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsTerminal(t *testing.T) {
	t.Run("PTY", func(t *testing.T) {
		pty, tty, err := pty.Open()
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = pty.Close()
			_ = tty.Close()
		})

		assert.True(t, IsTerminal(pty))
	})

	t.Run("other writers", func(t *testing.T) {
		assert.False(t, IsTerminal(&bytes.Buffer{}))
	})

	t.Run("file", func(t *testing.T) {
		tmp := t.TempDir()
		f, err := os.Create(filepath.Join(tmp, "test"))
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = f.Close()
		})

		assert.False(t, IsTerminal(f))
	})
}
