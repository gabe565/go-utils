package cobrax

import (
	"runtime"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestBuildUserAgent(t *testing.T) {
	osArch := " (" + runtime.GOOS + "/" + runtime.GOARCH + ")"

	type args struct {
		version, commit string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no version or commit", args{}, "test" + osArch},
		{"version", args{version: "1.0.0"}, "test/v1.0.0" + osArch},
		{"commit", args{commit: "deadbeef"}, "test/deadbeef" + osArch},
		{"version and commit", args{version: "1.0.0", commit: "deadbeef"}, "test/v1.0.0-deadbeef" + osArch},
		{"commit has changes", args{commit: "*deadbeef"}, "test/deadbeef" + osArch},
		{"non-numeric version", args{version: "beta"}, "test/beta" + osArch},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Use: "test",
				Annotations: map[string]string{
					VersionKey: tt.args.version,
					CommitKey:  tt.args.commit,
				},
			}
			assert.Equal(t, tt.want, BuildUserAgent(cmd))
		})
	}
}

func TestGetCommit(t *testing.T) {
	cmd := &cobra.Command{
		Annotations: map[string]string{CommitKey: "deadbeef"},
	}
	assert.Equal(t, "deadbeef", GetCommit(cmd))
}

func TestGetVersion(t *testing.T) {
	cmd := &cobra.Command{
		Annotations: map[string]string{VersionKey: "1.0.0"},
	}
	assert.Equal(t, "1.0.0", GetVersion(cmd))
}
