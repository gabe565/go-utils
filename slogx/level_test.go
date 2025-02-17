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

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{"trace", LevelTrace, "trace"},
		{"debug", LevelDebug, "debug"},
		{"info", LevelInfo, "info"},
		{"warn", LevelWarn, "warn"},
		{"error", LevelError, "error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.String())
		})
	}
}

func TestLevel_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Level
		wantErr require.ErrorAssertionFunc
	}{
		{"trace", args{[]byte("trace")}, LevelTrace, require.NoError},
		{"trace+1", args{[]byte("trace+1")}, LevelDebug, require.NoError},
		{"debug", args{[]byte("debug")}, LevelDebug, require.NoError},
		{"debug-1", args{[]byte("debug-1")}, LevelTrace, require.NoError},
		{"info", args{[]byte("info")}, LevelInfo, require.NoError},
		{"warn", args{[]byte("warn")}, LevelWarn, require.NoError},
		{"error", args{[]byte("error")}, LevelError, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l Level
			tt.wantErr(t, l.UnmarshalText(tt.args.b))
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLevel_Interfaces(t *testing.T) {
	l := LevelInfo
	assert.Implements(t, (*fmt.Stringer)(nil), &l)
	assert.Implements(t, (*encoding.TextMarshaler)(nil), &l)
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), &l)
	assert.Implements(t, (*json.Marshaler)(nil), &l)
	assert.Implements(t, (*json.Unmarshaler)(nil), &l)
	assert.Implements(t, (*pflag.Value)(nil), &l)
}
