package message

import (
	"fmt"
	"os"
	"strings"
)

const (
	usageCmdVerbNameF = "$ %s %s %s <message-name>"
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
	intro := `💬 MESSAGES:
The frontend and backend are separate go threads that communicate using messages.
Message definitions are at "shared/message/".
The backend message handlers are at "backend/txrx/"
The frontend message handlers are in each panel-group's messenger.go file.
A message name is upper camel case so a valid message name might be "AddContact".

You can add a message name with:     "$ kickfyne message add <message-name>"
You can remove a message name with:  "$ kickfyne message remove <message-name>"
You can list the message names with: "$ kickfyne message list"
You can view this information with:  "% kickfyne message help"

After you add a message name you will:
  1. Complete the message definition by adding your own custom elements to the default elements.
  2. Complete the message's backend handler.
  3. Add frontend message handlers to each panel-group-messenger that needs to handle the message.

MANAGING MESSAGES:
`
	usage = intro + commands + "\n"
	return
}
