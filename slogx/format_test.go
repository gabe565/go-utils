package slogx

import (
	"encoding"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat_String(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want string
	}{
		{"auto", FormatAuto, "auto"},
		{"color", FormatColor, "color"},
		{"plain", FormatPlain, "plain"},
		{"json", FormatJSON, "json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.f.String())
		})
	}
}

func TestFormat_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Format
		wantErr require.ErrorAssertionFunc
	}{
		{"auto", args{[]byte("auto")}, FormatAuto, require.NoError},
		{"color", args{[]byte("color")}, FormatColor, require.NoError},
		{"plain", args{[]byte("plain")}, FormatPlain, require.NoError},
		{"json", args{[]byte("json")}, FormatJSON, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Format
			tt.wantErr(t, f.UnmarshalText(tt.args.text))
			assert.Equal(t, tt.want, f)
		})
	}
}

func TestFormat_Interfaces(t *testing.T) {
	f := FormatAuto
	assert.Implements(t, (*fmt.Stringer)(nil), &f)
	assert.Implements(t, (*encoding.TextMarshaler)(nil), &f)
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), &f)
	assert.Implements(t, (*json.Marshaler)(nil), &f)
	assert.Implements(t, (*json.Unmarshaler)(nil), &f)
	assert.Implements(t, (*pflag.Value)(nil), &f)
}
