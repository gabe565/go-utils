package termx

import (
	"os"

	"github.com/mattn/go-isatty"
)

// IsTerminal returns whether the given io.Writer is a terminal.
func IsTerminal(v any) bool {
	if f, ok := v.(*os.File); ok {
		return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
	}
	return false
}
