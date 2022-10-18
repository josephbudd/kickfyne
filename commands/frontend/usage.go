package frontend

import (
	"fmt"
	"os"
	"strings"
)

const (
	newLine                = "\n"
	screenPackageNameParam = "«screen-package-name»"
	panelNameParam         = "«panel-name»"

	usage3F = "$ %s %s %s"
	usage4F = "$ %s %s %s %s"
	usage5F = "$ %s %s %s %s %s"
)

var (
	UsageScreenAdd          = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbAdd, screenPackageNameParam)
	UsageScreenAddAccordion = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbAddAccordion, screenPackageNameParam)
	UsageScreenAddAppTabs   = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbAddAppTabs, screenPackageNameParam)
	UsageScreenAddDocTabs   = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbAddDocTabs, screenPackageNameParam)
	usageScreenHelp         = fmt.Sprintf(usage3F, os.Args[0], CmdScreen, subCmdHelp)
	UsageScreenLanding      = fmt.Sprintf(usage3F, os.Args[0], CmdScreen, subCmdLanding)
	usageScreenList         = fmt.Sprintf(usage3F, os.Args[0], CmdScreen, verbList)
	UsageScreenRemove       = fmt.Sprintf(usage4F, os.Args[0], CmdScreen, verbRemove, screenPackageNameParam)

	UsagePanelAdd    = fmt.Sprintf(usage5F, os.Args[0], CmdPanel, verbAdd, screenPackageNameParam, panelNameParam)
	usagePanelHelp   = fmt.Sprintf(usage3F, os.Args[0], CmdPanel, subCmdHelp)
	usagePanelList   = fmt.Sprintf(usage4F, os.Args[0], CmdPanel, verbList, screenPackageNameParam)
	UsagePanelRemove = fmt.Sprintf(usage5F, os.Args[0], CmdPanel, verbRemove, screenPackageNameParam, panelNameParam)
)

func UsageScreen() (usage string) {
	commands := []string{
		UsageScreenAdd,
		UsageScreenAddAccordion, UsageScreenAddAppTabs, UsageScreenAddDocTabs,
		UsageScreenRemove,
		usageScreenList,
		UsageScreenLanding,
		usageScreenHelp,
	}
	usage = `📺 MANAGING SCREENS:
` + strings.Join(commands, newLine) + `
After a screen is added:
1. A search for KICKFYNE TODO will reveal instructions for proper developement and management of the screen operation.
`
	return
}

func UsagePanel() (usage string) {
	commands := []string{
		UsagePanelAdd,
		UsagePanelRemove,
		usagePanelList,
		usagePanelHelp,
	}
	usage = `📊 MANAGING SCREEN PANELS:
` + strings.Join(commands, newLine) + `
After a panel is added:
1. A search for KICKFYNE TODO will reveal instructions for proper developement and management of the panel operation.
`
	return
}

func Usage() (usage string) {
	usage = `👀 THE front-end:
The front-end's user interface is built with screen packages. A screen package is composed of panels and a messenger. The panels are how the screen package interacts with the user, displaying content and accepting user input. The messenger is how the screen package communicates with the back-end, sending and receiving messages.

For example a screen might have a panel where the user selects a record to edit and then another panel for editing the record. Meanwhile the messenger is communicating with the back-end getting records for selection, a record for editing, and sending the edited record to the back-end for storage and switching panels.

Alternately, the above screen with two panels could be done with two screens. One screen with only a select panel and the other screen with only an edit panel.

` + UsageScreen() + newLine + UsagePanel()
	return
}
