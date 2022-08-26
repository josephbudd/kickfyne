package record

import (
	"fmt"
	"os"
	"strings"
)

const (
	usageCmdVerbNameF = "$ %s %s %s <record-name>"
	usageCmdVerbF     = "$ %s %s %s"
	instructionsF     = " (For %s instructions.)"
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
	usageHelp = fmt.Sprintf(usageCmdVerbF, os.Args[0], Cmd, verbHelp) + fmt.Sprintf(instructionsF, Cmd)
}

func Usage() (usage string) {
	lines := []string{
		usageAdd,
		usageRemove,
		usageList,
		usageHelp,
	}
	commands := strings.Join(lines, "\n")
	intro := `💾 RECORDS:
When you create a record you also create it's store interface and implementation.
Record definitions are at shared/store/record/.
Record store interfaces are at shared/store/storer/.
Record store implementations are at shared/store/storing/.
A record name is upper camel case so a valid record name might be "Contact".

You can add a record name with:     "$ kickfyne record add <record-name>"
You can remove a record name with:  "$ kickfyne record remove <record-name>"
You can list the record names with: "$ kickfyne record list"
You can view this information with: "% kickfyne record help"

After you add a record:
  1. You will complete the record definition by adding your own custom elements.
  2. You may also modify the record's store interface and implementation.
     The records are by default, stored locally in easy to read yaml files.
     You are free to change that by modifying a record's store interface and implementation. 

MANAGING RECORDS AND STORES:
`
	usage = intro + commands + "\n"
	return
}
