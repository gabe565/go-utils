package cobrax

import (
	"runtime/debug"

	"gabe565.com/utils/httpx"
	"gabe565.com/utils/versionx"
	"github.com/spf13/cobra"
)

// GetVersion gets the raw version set by WithVersion.
func GetVersion(cmd *cobra.Command) string {
	if v, ok := cmd.Root().Annotations[VersionKey]; ok {
		return v
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return ""
}

// GetCommit gets the commit set by WithVersion.
func GetCommit(cmd *cobra.Command) string {
	if v, ok := cmd.Root().Annotations[CommitKey]; ok {
		return v
	}
	return versionx.CommitFromVCS().Short()
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
