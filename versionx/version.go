package versionx

func New(version string) Version {
	return Version{
		Version: version,
		Commit:  CommitFromVCS(),
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
