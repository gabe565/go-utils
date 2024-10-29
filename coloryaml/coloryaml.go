package coloryaml

import (
	"strings"

	"github.com/goccy/go-yaml/lexer"
)

// Colorize takes a serialized YAML string and returns it with ANSI color codes.
func Colorize(s string) string {
	// https://github.com/mikefarah/yq/blob/v4.43.1/pkg/yqlib/color_print.go
	tokens := lexer.Tokenize(s)
	return Printer().PrintTokens(tokens)
}
