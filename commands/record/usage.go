package record

import (
	"fmt"
	"os"
	"strings"
)

const (
	usageCmdVerbNameF = "$ %s %s %s «record-name»"
	usageCmdVerbF     = "$ %s %s %s"
)

var (
	usageAdd    string
	usageRemove string
	usageList   string
	usageHelp   string
)

func init() {
	usageAdd = fmt.Sprintf(usageCmdVerbNameF, os.Args[0], Cmd, verbAdd)
	usageRemove = fmt.Sprintf(usageCmdVerbNameF, os.Args[0], Cmd, verbRemove)
	usageList = fmt.Sprintf(usageCmdVerbF, os.Args[0], Cmd, verbList)
	usageHelp = fmt.Sprintf(usageCmdVerbF, os.Args[0], Cmd, verbHelp)
}

func Usage() (usage string) {
	commands := []string{
		usageAdd,
		usageRemove,
		usageList,
		usageHelp,
	}
	usage = `💾 RECORDS:

When a record is added, the record's store interface and implementation are also added.

MANAGING RECORDS AND STORES:
` + strings.Join(commands, "\n") + `

* Record definitions are at shared/store/record/.
* Record store interfaces are at shared/store/storer/.
* Record store implementations are at shared/store/storing/.
* A record name is upper camel case so a valid record name might be "Contact".

After a record is added:
1. A search for KICKFYNE TODO will reveal instructions for proper developement and management of the record and store operation.
   1. The record definition needs to be completed so that the message can contain useful information.
   2. The record's store interface and implementation may need modification. The records are by default, stored locally in easy to read yaml files. If that is not how the application is to handle data then modifications can be made to a record's store interface and implementation.

`
	return
}
