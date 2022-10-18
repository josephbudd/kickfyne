package message

import (
	"fmt"
	"os"
	"strings"
)

const (
	usageCmdVerbNameF = "$ %s %s %s «message-name»"
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
	usage = `💬 MESSAGES:
The front-end and back-end are separate go threads that communicate using messages.

MANAGING MESSAGES:
` + strings.Join(commands, "\n") + `
* Message definitions are at "shared/message/".
* The back-end message handlers are at "backend/txrx/"
* The front-end message handlers are in each screen's messageHandler.go file.
* A message name is upper camel case so a valid message name might be "AddContact".

After a message is added:
1. A search for KICKFYNE TODO will reveal instructions for proper developement and management of the message and handler operation.
   1. The message definition needs to be completed so that the message can contain useful information.
   2. The message's back-end handler needs functionality added.
2. Some front-end message handlers will need to handle sending and or receiving the message.
`
	return
}
