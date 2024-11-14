package colorx

import (
	"encoding"
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

var (
	_ encoding.TextMarshaler   = &Hex{}
	_ encoding.TextUnmarshaler = &Hex{}
	_ fmt.Stringer             = &Hex{}
)

// Hex is a helper struct that wraps a color and satisfies
// encoding.TextMarshaler, encoding.TextUnmarshaler, and fmt.Stringer.
type Hex struct {
	color.Color
}

// MarshalText encodes a color as hex.
func (h *Hex) MarshalText() ([]byte, error) {
	s := FormatHex(h.Color)
	return []byte(s), nil
}

// UnmarshalText parses the given hex code.
func (h *Hex) UnmarshalText(text []byte) error {
	c, err := ParseHex(string(text))
	if err != nil {
		return err
	}

	h.Color = c
	return nil
}

// MarshalText encodes a color as hex.
func (h *Hex) String() string {
	return FormatHex(h.Color)
}

var ErrInvalidLength = errors.New("hex code should be 4 or 7 characters")

// ParseHex parses the given hex code as a color.NRGBA
//
//nolint:gosec
func ParseHex(text string) (color.Color, error) {
	var c color.NRGBA
	text = strings.TrimPrefix(text, "#")

	switch len(text) {
	case 3, 4:
		parsed, err := strconv.ParseUint(text, 16, 16)
		if err != nil {
			return c, err
		}

		if len(text) == 4 {
			c.A = uint8(parsed & 0xF)
			c.A |= c.A << 4
			parsed >>= 4
		} else {
			c.A = 0xFF
		}
		c.B = uint8(parsed & 0xF)
		c.B |= c.B << 4
		parsed >>= 4
		c.G = uint8(parsed & 0xF)
		c.G |= c.G << 4
		parsed >>= 4
		c.R = uint8(parsed & 0xF)
		c.R |= c.R << 4
	case 6, 8:
		parsed, err := strconv.ParseUint(text, 16, 32)
		if err != nil {
			return c, err
		}

		if len(text) == 8 {
			c.A = uint8(parsed & 0xFF)
			parsed >>= 8
		} else {
			c.A = 0xFF
		}
		c.B = uint8(parsed & 0xFF)
		parsed >>= 8
		c.G = uint8(parsed & 0xFF)
		parsed >>= 8
		c.R = uint8(parsed & 0xFF)
	default:
		return c, ErrInvalidLength
	}

	if c.R == c.G && c.G == c.B && c.A == 0xFF {
		y := uint16(c.R)
		y |= y << 8
		return color.Gray16{Y: y}, nil
	}
	return c, nil
}

// FormatHex formats the given color.Color as a hex code
//
//nolint:gosec
func FormatHex(c color.Color) string {
	var r, g, b, a uint8
	switch c := c.(type) {
	case color.NRGBA:
		r, g, b, a = c.R, c.G, c.B, c.A
	default:
		r32, g32, b32, a32 := c.RGBA()
		r, g, b, a = uint8(r32&0xFF), uint8(g32&0xFF), uint8(b32&0xFF), uint8(a32&0xFF)
	}

	shorthand := r>>4 == r&0xF && g>>4 == g&0xF && b>>4 == b&0xF
	skipAlpha := a&0xFF == 0xFF
	if shorthand {
		if skipAlpha {
			return fmt.Sprintf("#%x%x%x", r&0xF, g&0xF, b&0xF)
		}
		return fmt.Sprintf("#%x%x%x%x", r&0xF, g&0xF, b&0xF, a&0xF)
	}
	if skipAlpha {
		return fmt.Sprintf("#%02x%02x%02x", r, g, b)
	}
	return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}
