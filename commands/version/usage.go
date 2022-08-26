package version

import (
	"fmt"
	"os"
)

const (
	Cmd          = "version"
	usageCmdCmdF = "$ %s %s"
)

var (
	usageVersion string
)

func init() {
	usageVersion = fmt.Sprintf(usageCmdCmdF, os.Args[0], Cmd)
}

func Usage() (usage string) {
	usage = usageVersion
	return
}
