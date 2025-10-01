package bytefmt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// NewDecoder creates a new Decoder.
func NewDecoder() *Decoder {
	return &Decoder{}
}

// A Decoder decodes strings to byte values.
type Decoder struct{}

func split(text string) (float64, string, error) {
	suffix := strings.TrimLeftFunc(text, func(r rune) bool {
		return ('0' <= r && r <= '9') || r == ' ' || r == '-' || r == '.'
	})

	bytesStr := strings.TrimSuffix(text[:len(text)-len(suffix)], " ")
	bytes, err := strconv.ParseFloat(bytesStr, 64)
	if err != nil {
		return 0, "", err
	}

	return bytes, suffix, nil
}

// Decode parses human-readable bytes string to bytes integer.
//
// For example, "6 GiB" ("6 Gi" is also valid) will return 6442450944, and
// "6 GB" ("6 G" is also valid) will return 6000000000.
// The input case-insensitive, and the space is optional,
// so "6GiB" or "6GB" would produce the same output.
func (b *Decoder) Decode(text string) (int64, error) {
	bytes, suffix, err := split(text)
	if err != nil {
		return 0, err
	}

	if val, err := b.decodeBinary(bytes, suffix); err == nil {
		return val, nil
	}

	return b.decodeDecimal(bytes, suffix)
}

var ErrUnknownSuffix = errors.New("unknown suffix")

// DecodeBinary parses human-readable bytes string to bytes integer.
//
// For example, "6 GiB" ("6 Gi" is also valid) will return 6442450944.
// The input case-insensitive, and the space is optional,
// so "6GiB" would produce the same output.
func (b *Decoder) DecodeBinary(text string) (int64, error) {
	bytes, suffix, err := split(text)
	if err != nil {
		return 0, err
	}

	return b.decodeBinary(bytes, suffix)
}

func (b *Decoder) decodeBinary(bytes float64, suffix string) (int64, error) {
	switch strings.ToUpper(suffix) {
	case "KI", "KIB":
		return int64(bytes * KiB), nil
	case "MI", "MIB":
		return int64(bytes * MiB), nil
	case "GI", "GIB":
		return int64(bytes * GiB), nil
	case "TI", "TIB":
		return int64(bytes * TiB), nil
	case "PI", "PIB":
		return int64(bytes * PiB), nil
	case "EI", "EIB":
		return int64(bytes * EiB), nil
	case "":
		return int64(bytes), nil
	default:
		return 0, fmt.Errorf("%w: %q", ErrUnknownSuffix, suffix)
	}
}

// DecodeDecimal parses human-readable bytes string to bytes integer.
//
// For example, "6 GB" ("6 G" is also valid) will return 6000000000.
// The input case-insensitive, and the space is optional,
// so "6GB" would produce the same output.
func (b *Decoder) DecodeDecimal(text string) (int64, error) {
	bytes, suffix, err := split(text)
	if err != nil {
		return 0, err
	}

	return b.decodeDecimal(bytes, suffix)
}

func (b *Decoder) decodeDecimal(bytes float64, suffix string) (int64, error) {
	switch strings.ToUpper(suffix) {
	case "K", "KB":
		return int64(bytes * KB), nil
	case "M", "MB":
		return int64(bytes * MB), nil
	case "G", "GB":
		return int64(bytes * GB), nil
	case "T", "TB":
		return int64(bytes * TB), nil
	case "P", "PB":
		return int64(bytes * PB), nil
	case "E", "EB":
		return int64(bytes * EB), nil
	case "B", "":
		return int64(bytes), nil
	default:
		return 0, fmt.Errorf("%w: %q", ErrUnknownSuffix, suffix)
	}
}
