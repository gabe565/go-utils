package coloryaml

import (
	"io"
	"strings"

	"gabe565.com/utils/termx"
	"github.com/goccy/go-yaml/lexer"
)

// Colorize takes a serialized YAML string and returns it with ANSI color codes.
func Colorize(s string) string {
	// https://github.com/mikefarah/yq/blob/v4.43.1/pkg/yqlib/color_print.go
	tokens := lexer.Tokenize(s)
	return Printer().PrintTokens(tokens)
}

// WriteString writes a YAML string to the given io.Writer.
// The written data will be colorized if termx.IsColor returns true,
// and the output will always end with a new line character.
func WriteString(w io.Writer, s string) (int, error) {
	if termx.IsColor(w) {
		s = strings.TrimSuffix(s, "\n")
		s = Colorize(s)
	}

	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}

	return io.WriteString(w, s)
}
