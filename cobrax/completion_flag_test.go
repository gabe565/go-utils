package cobrax

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenCompletion(t *testing.T) {
	type args struct {
		shell Shell
	}
	tests := []struct {
		name    string
		args    args
		wantErr require.ErrorAssertionFunc
	}{
		{"bash", args{Bash}, require.NoError},
		{"zsh", args{Zsh}, require.NoError},
		{"fish", args{Fish}, require.NoError},
		{"powershell", args{PowerShell}, require.NoError},
		{"other", args{"other"}, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			var buf strings.Builder
			cmd.SetOut(&buf)
			err := GenCompletion(cmd, tt.args.shell)
			tt.wantErr(t, err)
			if err == nil {
				assert.NotZero(t, buf.Len())
			}
		})
	}
}

func TestRegisterCompletionFlag(t *testing.T) {
	cmd := &cobra.Command{}
	require.NoError(t, RegisterCompletionFlag(cmd))
	f := cmd.Flags().Lookup(FlagCompletion)
	require.NotNil(t, f)

	_, registered := cmd.GetFlagCompletionFunc(FlagCompletion)
	assert.True(t, registered)
}
