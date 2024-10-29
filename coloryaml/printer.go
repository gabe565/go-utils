package coloryaml

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml/printer"
)

const escape = "\x1b"

func format(attr color.Attribute) string {
	return escape + "[" + strconv.Itoa(int(attr)) + "m"
}

// Printer returns a printer.Printer pre-configured to add ANSI color codes to YAML.
func Printer() *printer.Printer {
	return &printer.Printer{
		MapKey: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgCyan),
				Suffix: format(color.Reset),
			}
		},
		Anchor: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiYellow),
				Suffix: format(color.Reset),
			}
		},
		Alias: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiYellow),
				Suffix: format(color.Reset),
			}
		},
		Bool: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiMagenta),
				Suffix: format(color.Reset),
			}
		},
		String: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgGreen),
				Suffix: format(color.Reset),
			}
		},
		Number: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiMagenta),
				Suffix: format(color.Reset),
			}
		},
		Comment: func() *printer.Property {
			return &printer.Property{
				Prefix: format(color.FgHiBlack),
				Suffix: format(color.Reset),
			}
		},
	}
}
