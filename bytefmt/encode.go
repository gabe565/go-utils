package bytefmt

import (
	"strconv"
	"strings"
)

const (
	DefaultPrecision   = 2
	DefaultTrimIntZero = true
	DefaultUseSpace    = true
)

// NewEncoder creates a new Encoder.
func NewEncoder() *Encoder {
	return &Encoder{
		precision:      DefaultPrecision,
		trimIntDecimal: DefaultTrimIntZero,
		useSpace:       DefaultUseSpace,
	}
}

// A Encoder encodes byte values to strings.
type Encoder struct {
	precision      int
	trimIntDecimal bool
	useSpace       bool
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
	b.useSpace = useSpace
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
	case valInt >= EiB:
		valFloat /= EiB
		multiple = "EiB"
	case valInt >= PiB:
		valFloat /= PiB
		multiple = "PiB"
	case valInt >= TiB:
		valFloat /= TiB
		multiple = "TiB"
	case valInt >= GiB:
		valFloat /= GiB
		multiple = "GiB"
	case valInt >= MiB:
		valFloat /= MiB
		multiple = "MiB"
	case valInt >= KiB:
		valFloat /= KiB
		multiple = "KiB"
	case valInt == 0:
		return "0"
	default:
		output := strconv.FormatInt(valInt, 10)
		if b.useSpace {
			return output + " B"
		}
		return output + "B"
	}
	if b.useSpace {
		multiple = " " + multiple
	}
	output := strconv.FormatFloat(valFloat, 'f', b.precision, 64) + multiple
	if b.trimIntDecimal {
		output = strings.Replace(output, ".00 ", " ", 1)
	}
	return output
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
	case valInt >= EB:
		valFloat /= EB
		multiple = "EB"
	case valInt >= PB:
		valFloat /= PB
		multiple = "PB"
	case valInt >= TB:
		valFloat /= TB
		multiple = "TB"
	case valInt >= GB:
		valFloat /= GB
		multiple = "GB"
	case valInt >= MB:
		valFloat /= MB
		multiple = "MB"
	case valInt >= KB:
		valFloat /= KB
		multiple = "KB"
	case valInt == 0:
		return "0"
	default:
		output := strconv.FormatInt(valInt, 10)
		if b.useSpace {
			return output + " B"
		}
		return output + "B"
	}
	if b.useSpace {
		multiple = " " + multiple
	}
	output := strconv.FormatFloat(valFloat, 'f', b.precision, 64) + multiple
	if b.trimIntDecimal {
		output = strings.Replace(output, ".00 ", " ", 1)
	}
	return output
}
