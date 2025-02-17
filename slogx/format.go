package slogx

import (
	"bytes"
	"errors"
	"strconv"
)

type Format uint8

const (
	FormatAuto Format = iota
	FormatColor
	FormatPlain
	FormatJSON
)

func FormatStrings() []string {
	return []string{
		FormatAuto.String(),
		FormatColor.String(),
		FormatPlain.String(),
		FormatJSON.String(),
	}
}

var ErrUnknownFormat = errors.New("unknown format")

func (f *Format) UnmarshalText(text []byte) error {
	switch {
	case bytes.EqualFold(text, []byte("auto")):
		*f = FormatAuto
	case bytes.EqualFold(text, []byte("color")):
		*f = FormatColor
	case bytes.EqualFold(text, []byte("plain")):
		*f = FormatPlain
	case bytes.EqualFold(text, []byte("json")):
		*f = FormatJSON
	default:
		return ErrUnknownFormat
	}
	return nil
}

func (f Format) String() string {
	switch f {
	case FormatAuto:
		return "auto"
	case FormatColor:
		return "color"
	case FormatPlain:
		return "plain"
	case FormatJSON:
		return "json"
	default:
		return "slogx.Format(" + strconv.Itoa(int(f)) + ")"
	}
}

func (f Format) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

func (f Format) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, f.String()), nil
}

func (f *Format) UnmarshalJSON(i []byte) error {
	s, err := strconv.Unquote(string(i))
	if err != nil {
		return err
	}
	return f.Set(s)
}

func (f *Format) Set(s string) error {
	return f.UnmarshalText([]byte(s))
}

func (f Format) Type() string {
	return "string"
}
