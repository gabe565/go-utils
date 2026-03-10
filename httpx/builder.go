package httpx

import (
	"runtime"
	"strings"

	"gabe565.com/utils/versionx"
)

// BuildUserAgent builds an opinionated User-Agent string.
// If commit is empty, it will be set to the short commit hash of the current
// commit.
//
// Example output: `example/v1.0.0-deadbeef (linux/amd64)`.
func BuildUserAgent(name, version, commit string) string {
	if commit == "" {
		commit = versionx.CommitFromVCS().Short()
	}
	commit = strings.TrimPrefix(commit, "*")

	var ua strings.Builder
	ua.Grow(len(name) + len(version) + len(commit) + len(runtime.GOOS) + len(runtime.GOARCH) + 10)
	ua.WriteString(name)

	if version != "" {
		ua.WriteByte('/')
		if '0' <= version[0] && version[0] <= '9' {
			ua.WriteByte('v')
		}
		ua.WriteString(version)
		if commit != "" {
			ua.WriteByte('-')
			ua.WriteString(commit)
		}
	} else if commit != "" {
		ua.WriteByte('/')
		ua.WriteString(commit)
	}

	ua.WriteString(" (")
	ua.WriteString(runtime.GOOS)
	ua.WriteByte('/')
	ua.WriteString(runtime.GOARCH)
	ua.WriteByte(')')

	return ua.String()
}
