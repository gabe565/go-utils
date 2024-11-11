package cobrax

import (
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

func buildVersion(version string) (string, string) {
	var commit string
	var modified bool
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				commit = setting.Value
			case "vcs.modified":
				if setting.Value == "true" {
					modified = true
				}
			}
		}
	}

	if commit != "" {
		if len(commit) > 8 {
			commit = commit[:8]
		}
		if modified {
			commit = "*" + commit
		}
		version += " (" + commit + ")"
	}
	return version, commit
}

// GetVersion gets the raw version set by WithVersion
func GetVersion(cmd *cobra.Command) string {
	return cmd.Root().Annotations[VersionKey]
}

// GetCommit gets the commit set by WithVersion
func GetCommit(cmd *cobra.Command) string {
	return cmd.Root().Annotations[CommitKey]
}

// BuildUserAgent generates a value to be used with httpx.UserAgentTransport.
//
// Example output: `example/v1.0.0-deadbeef (linux/amd64)`
func BuildUserAgent(cmd *cobra.Command) string {
	root := cmd.Root()
	ua := root.Name()
	commit := strings.TrimPrefix(GetCommit(root), "*")
	if version := GetVersion(root); version != "" {
		ua += "/v" + version
		if commit != "" {
			ua += "-" + commit
		}
	} else if commit != "" {
		ua += "/" + commit
	}
	ua += " (" + runtime.GOOS + "/" + runtime.GOARCH + ")"
	return ua
}
