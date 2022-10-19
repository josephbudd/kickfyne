package version

import (
	"fmt"
	"os"
)

const (
	versionNewBreakingAPI  = 0
	versionAPIAddedFeature = 4
	versionAPIBugFix       = 1
)

// V returns the version.
func V() (version string) {
	version = fmt.Sprintf("%s v%d.%d.%d", os.Args[0], versionNewBreakingAPI, versionAPIAddedFeature, versionAPIBugFix)
	return
}
