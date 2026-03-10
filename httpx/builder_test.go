package httpx

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUserAgent(t *testing.T) {
	osArch := " (" + runtime.GOOS + "/" + runtime.GOARCH + ")"

	type args struct {
		name    string
		version string
		commit  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "name only", args: args{name: "test"}, want: "test" + osArch},
		{name: "numeric version", args: args{name: "test", version: "1.0.0"}, want: "test/v1.0.0" + osArch},
		{name: "non-numeric version", args: args{name: "test", version: "beta"}, want: "test/beta" + osArch},
		{name: "commit only", args: args{name: "test", commit: "deadbeef"}, want: "test/deadbeef" + osArch},
		{
			name: "version and commit",
			args: args{name: "test", version: "1.0.0", commit: "deadbeef"},
			want: "test/v1.0.0-deadbeef" + osArch,
		},
		{
			name: "modified commit is trimmed",
			args: args{name: "test", commit: "*deadbeef"},
			want: "test/deadbeef" + osArch,
		},
		{
			name: "long commit is truncated",
			args: args{name: "test", commit: "deadbeef"},
			want: "test/deadbeef" + osArch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildUserAgent(tt.args.name, tt.args.version, tt.args.commit)
			assert.Equal(t, tt.want, got)
		})
	}
}
