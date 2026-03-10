package versionx

import "runtime/debug"

func CommitFromVCS() Commit {
	var commit Commit
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				commit.SHA = setting.Value
			case "vcs.modified":
				if setting.Value == "true" {
					commit.Modified = true
				}
			}
		}
	}
	return commit
}

type Commit struct {
	SHA      string
	Modified bool
}

func (c Commit) Long() string {
	if c.Modified && c.SHA != "" {
		return "*" + c.SHA
	}
	return c.SHA
}

func (c Commit) Short() string {
	sha := c.SHA
	if len(sha) > 8 {
		sha = sha[:8]
	}
	if c.Modified && sha != "" {
		sha = "*" + sha
	}
	return sha
}
