package colorx

import (
	"encoding"
	"fmt"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHex(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), &Hex{})
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), &Hex{})
	assert.Implements(t, (*fmt.Stringer)(nil), &Hex{})
}

func TestParseHex(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    color.Color
		wantErr require.ErrorAssertionFunc
	}{
		{"white", args{"#fff"}, color.White, require.NoError},
		{"black", args{"#000"}, color.Black, require.NoError},
		{"red", args{"#f00"}, color.NRGBA{R: 0xFF, A: 0xFF}, require.NoError},
		{"green", args{"#0f0"}, color.NRGBA{G: 0xFF, A: 0xFF}, require.NoError},
		{"blue", args{"#00f"}, color.NRGBA{B: 0xFF, A: 0xFF}, require.NoError},
		{"blue gray", args{"#607d8b"}, color.NRGBA{R: 0x60, G: 0x7D, B: 0x8B, A: 0xFF}, require.NoError},
		{"missing prefix", args{"fff"}, color.White, require.NoError},
		{"short alpha", args{"#f00e"}, color.NRGBA{R: 0xFF, A: 0xEE}, require.NoError},
		{"long alpha", args{"#00ff00ab"}, color.NRGBA{G: 0xFF, A: 0xAB}, require.NoError},
		{"too long", args{"#fffffffff"}, nil, require.Error},
		{"too short", args{"#fffff"}, nil, require.Error},
		{"invalid character", args{"#ffz"}, nil, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := ParseHex(tt.args.text)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, c)
		})
	}
}

func TestFormatHex(t *testing.T) {
	type fields struct {
		c color.Color
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"white", fields{color.White}, "#fff"},
		{"black", fields{color.Black}, "#000"},
		{"red", fields{color.NRGBA{R: 0xFF, A: 0xFF}}, "#f00"},
		{"green", fields{color.NRGBA{G: 0xFF, A: 0xFF}}, "#0f0"},
		{"blue", fields{color.NRGBA{B: 0xFF, A: 0xFF}}, "#00f"},
		{"blue-gray", fields{color.NRGBA{R: 0x60, G: 0x7D, B: 0x8B, A: 0xFF}}, "#607d8b"},
		{"increment", fields{color.NRGBA{R: 1, G: 2, B: 3, A: 0xFF}}, "#010203"},
		{"alpha", fields{color.NRGBA{B: 0xFF, A: 0xAA}}, "#00fa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FormatHex(tt.fields.c))
		})
	}
}
