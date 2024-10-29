package termx

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// IsTerminal returns whether the given io.Writer is a terminal.
func IsTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
	}
	return false
}

// IsColor returns whether color should be enabled. It will only return true when:
//   - The NO_COLOR env is empty or unset
//   - The TERM env does not equal "dumb"
//   - The given io.Writer is a terminal
func IsColor(w io.Writer) bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	return IsTerminal(w)
}
