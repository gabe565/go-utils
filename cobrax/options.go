package cobrax

import "github.com/spf13/cobra"

type Option func(cmd *cobra.Command)

const (
	// VersionKey is the annotation key for the raw version.
	VersionKey = "go-utils_version"
	// CommitKey is the annotation key for the commit.
	CommitKey = "go-utils_commit"
)

// WithVersion takes a version string and sets the value to a cobra.Command.
// The current commit will be appended in parentheses, and the raw values
// will be added to the command's annotations.
func WithVersion(version string) Option {
	return func(cmd *cobra.Command) {
		if cmd.Annotations == nil {
			cmd.Annotations = make(map[string]string)
		}
		cmd.Annotations[VersionKey] = version
		cmd.Version, cmd.Annotations[CommitKey] = buildVersion(version)
		cmd.InitDefaultVersionFlag()
	}
}
