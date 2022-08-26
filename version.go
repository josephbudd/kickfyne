package main

import "fmt"

const (
	applicationName        = "kickfyne"
	versionNewBreakingAPI  = 0
	versionAPIAddedFeature = 1
	varsionAPIBugFix       = 0
)

var version = fmt.Sprintf("%s: version:%d.%d.%d", applicationName, versionNewBreakingAPI, versionAPIAddedFeature, varsionAPIBugFix)
