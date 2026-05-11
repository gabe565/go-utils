package versionx

import "runtime/debug"

func New(version string) Version {
	info, ok := debug.ReadBuildInfo()
	if ok && version == "" {
		version = info.Main.Version
	}
	return Version{
		Version: version,
		Commit:  commitFromVCS(info),
	}
}

type Version struct {
	Version string
	Commit  Commit
}

func (v Version) String() string {
	s := v.Version
	if v.Commit.SHA != "" {
		if s == "" {
			s = v.Commit.Short()
		} else {
			s += " (" + v.Commit.Short() + ")"
		}
	}
	return s
}
