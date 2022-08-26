package framework

import (
	"fmt"
	"os"
)

func Usage() (usage string) {
	usage = fmt.Sprintf("CREATE THE FRAMEWORK:\n$ %s %s", os.Args[0], Cmd)
	return
}
