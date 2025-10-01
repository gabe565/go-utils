package bytefmt

import (
	"strconv"
	"strings"
)

const (
	DefaultPrecision      = 2
	DefaultTrimIntDecimal = true
	DefaultNoSpace        = false
)

// NewEncoder creates a new Encoder.
func NewEncoder() *Encoder {
	return &Encoder{
		precision:      DefaultPrecision,
		trimIntDecimal: DefaultTrimIntDecimal,
		noSpace:        DefaultNoSpace,
	}
}

// A Encoder encodes byte values to strings.
type Encoder struct {
	precision      int
	trimIntDecimal bool
	noSpace        bool
}

// SetPrecision instructs the encoder to include the given number of decimal places.
//
// The special precision -1 uses the smallest number of digits necessary.
func (b *Encoder) SetPrecision(precision int) *Encoder {
	b.precision = precision
	return b
}

// SetTrimIntDecimal instructs the encoder to remove the decimal if the result is
// an integer (ends with ".00").
func (b *Encoder) SetTrimIntDecimal(trimIntZero bool) *Encoder {
	b.trimIntDecimal = trimIntZero
	return b
}

// SetUseSpace instructs the encoder to separate the value and suffix with a space.
func (b *Encoder) SetUseSpace(useSpace bool) *Encoder {
	b.noSpace = !useSpace
	return b
}

// Encode formats bytes integer to human-readable string according to IEC 60027.
//
// For example, 31323 bytes will return "30.59 KB".
//
// The smallest supported value is "1 B", so precision will be ignored.
// below "1 KiB" or "1 KB"
func (b *Encoder) Encode(val int64) string {
	return b.EncodeBinary(val)
}

// EncodeBinary formats bytes integer to human-readable string according to IEC 60027.
//
// For example, 31323 bytes will return "30.59 KB".
//
// The smallest supported value is "1 B", so precision will be ignored
// below "1 KB".
func (b *Encoder) EncodeBinary(valInt int64) string { //nolint:dupl
	var multiple string
	valFloat := float64(valInt)
	switch {
	case valInt == 0:
		return "0"
	case valInt < KiB:
		output := strconv.FormatInt(valInt, 10)
		if !b.noSpace {
			return output + " B"
		}
		return output + "B"
	case valInt < MiB:
		valFloat /= KiB
		multiple = "KiB"
	case valInt < GiB:
		valFloat /= MiB
		multiple = "MiB"
	case valInt < TiB:
		valFloat /= GiB
		multiple = "GiB"
	case valInt < PiB:
		valFloat /= TiB
		multiple = "TiB"
	case valInt < EiB:
		valFloat /= PiB
		multiple = "PiB"
	default:
		valFloat /= EiB
		multiple = "EiB"
	}
	if !b.noSpace {
		multiple = " " + multiple
	}
	output := strconv.FormatFloat(valFloat, 'f', b.precision, 64)
	if b.trimIntDecimal && b.precision != 0 {
		output = strings.TrimSuffix(output, ".00")
	}
	return output + multiple
}

// EncodeDecimal formats bytes integer to human-readable string according to SI international system of units.
//
// For example, 31323 bytes will return "31.32 KB".
//
// The smallest supported value is "1 B", so precision will be ignored
// below "1 KiB".
func (b *Encoder) EncodeDecimal(valInt int64) string { //nolint:dupl
	var multiple string
	valFloat := float64(valInt)
	switch {
	case valInt == 0:
		return "0"
	case valInt < KB:
		output := strconv.FormatInt(valInt, 10)
		if !b.noSpace {
			return output + " B"
		}
		return output + "B"
	case valInt < MB:
		valFloat /= KB
		multiple = "KB"
	case valInt < GB:
		valFloat /= MB
		multiple = "MB"
	case valInt < TB:
		valFloat /= GB
		multiple = "GB"
	case valInt < PB:
		valFloat /= TB
		multiple = "TB"
	case valInt < EB:
		valFloat /= PB
		multiple = "PB"
	default:
		valFloat /= EB
		multiple = "EB"
	}
	if !b.noSpace {
		multiple = " " + multiple
	}
	output := strconv.FormatFloat(valFloat, 'f', b.precision, 64)
	if b.trimIntDecimal && b.precision != 0 {
		output = strings.TrimSuffix(output, ".00")
	}
	return output + multiple
}
