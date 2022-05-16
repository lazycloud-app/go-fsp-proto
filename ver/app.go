package ver

import "fmt"

type (
	// AppVersion satisfies the requirements of https://semver.org/
	AppVersion struct {
		Major        uint
		MLabel       string
		Minor        uint
		Patch        uint
		ReleaseName  string
		ReleaseDate  string
		ReleaseSatus string
		Proto        Proto
	}
)

func (v AppVersion) String() string {
	return fmt.Sprintf(`%d.%d.%d`, v.Major, v.Minor, v.Patch)
}

func (v AppVersion) Full() string {
	return fmt.Sprintf(`%d.%d.%d-%s "%s" (%s) Branch:%s`, v.Major, v.Minor, v.Patch, v.MLabel, v.ReleaseName, v.ReleaseDate, v.ReleaseSatus)
}
