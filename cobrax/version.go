package cobrax

import (
	"gabe565.com/utils/httpx"
	"github.com/spf13/cobra"
)

// GetVersion gets the raw version set by WithVersion.
func GetVersion(cmd *cobra.Command) string {
	return cmd.Root().Annotations[VersionKey]
}

// GetCommit gets the commit set by WithVersion.
func GetCommit(cmd *cobra.Command) string {
	return cmd.Root().Annotations[CommitKey]
}

// BuildUserAgent generates a value to be used with httpx.UserAgentTransport.
//
// Example output: `example/v1.0.0-deadbeef (linux/amd64)`.
func BuildUserAgent(cmd *cobra.Command) string {
	return httpx.BuildUserAgent(
		cmd.Root().Name(),
		GetVersion(cmd),
		GetCommit(cmd),
	)
}
