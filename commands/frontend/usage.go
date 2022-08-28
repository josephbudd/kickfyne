package frontend

import (
	"fmt"
	"os"
	"strings"
)

const (
	newLine        = "\n"
	newLineNewLine = "\n\n"
	blankLine      = ""
	commentF       = " (For instructions on how to %s.)"
	commentSF      = " (For instructions on how to %s a %s.)"
	commentPF      = " (For instructions on how to %s a %s.)"

	usageCmdSubcmdVerbPathF     = "$ %s %s %s %s %s"
	usageCmdSubcmdVerbNameF     = "$ %s %s %s %s %s"
	usageCmdSubcmdVerbSubcmdF   = "$ %s %s %s %s %s"
	usageCmdSubcmdVerb2NamesF   = "$ %s %s %s %s %s %s"
	usageCmdSubcmdVerb3NamesF   = "$ %s %s %s %s %s %s %s"
	usageCmdSubcmdPathF         = "$ %s %s %s %s"
	usageCmdSubcmdVerbF         = "$ %s %s %s %s"
	usageCmdSubcmdVerbThisHelpF = "$ %s %s %s %s (This help.)"
	usageCmdSubcmdF             = "$ %s %s %s"
)

var (
	UsageCreate        = fmt.Sprintf(usageCmdSubcmdPathF, os.Args[0], Cmd, SubCmdCreate, "<path to create-buttons-yaml-file>")
	usageCreateHelp    = fmt.Sprintf(usageCmdSubcmdVerbF, os.Args[0], Cmd, SubCmdCreate, subCmdHelp) + fmt.Sprintf(commentF, SubCmdCreate)
	usageClear         = fmt.Sprintf(usageCmdSubcmdF, os.Args[0], Cmd, subCmdClear)
	UsageButtonAdd     = fmt.Sprintf(usageCmdSubcmdVerbPathF, os.Args[0], Cmd, SubCmdButton, verbAdd, "<path to add-button-yaml-file>")
	usageButtonAddHelp = fmt.Sprintf(usageCmdSubcmdVerbSubcmdF, os.Args[0], Cmd, SubCmdButton, verbAdd, subCmdHelp) + fmt.Sprintf(commentSF, verbAdd, SubCmdButton)
	usageButtonRemove  = fmt.Sprintf(usageCmdSubcmdVerbNameF, os.Args[0], Cmd, SubCmdButton, verbRemove, "<button-name>")
	usageButtonList    = fmt.Sprintf(usageCmdSubcmdVerbF, os.Args[0], Cmd, SubCmdButton, verbList)
	usageButtonHelp    = fmt.Sprintf(usageCmdSubcmdVerbThisHelpF, os.Args[0], Cmd, SubCmdButton, subCmdHelp)
	UsageTabAdd        = fmt.Sprintf(usageCmdSubcmdVerbPathF, os.Args[0], Cmd, SubCmdTab, verbAdd, "<path to add-tab-yaml-file>")
	usageTabAddHelp    = fmt.Sprintf(usageCmdSubcmdVerbSubcmdF, os.Args[0], Cmd, SubCmdTab, verbAdd, subCmdHelp) + fmt.Sprintf(commentSF, verbAdd, SubCmdTab)
	usageTabRemove     = fmt.Sprintf(usageCmdSubcmdVerb2NamesF, os.Args[0], Cmd, SubCmdTab, verbRemove, "<buton-name>", "<tab-name>")
	usageTabHelp       = fmt.Sprintf(usageCmdSubcmdVerbThisHelpF, os.Args[0], Cmd, SubCmdTab, subCmdHelp)
	UsagePanelAdd      = fmt.Sprintf(usageCmdSubcmdVerbPathF, os.Args[0], Cmd, SubCmdPanel, verbAdd, "<path to add-panel-yaml-file>")
	usagePanelAddHelp  = fmt.Sprintf(usageCmdSubcmdVerbSubcmdF, os.Args[0], Cmd, SubCmdPanel, verbAdd, subCmdHelp) + fmt.Sprintf(commentSF, verbAdd, SubCmdPanel)
	UsagePanelRemove   = fmt.Sprintf(usageCmdSubcmdVerb3NamesF, os.Args[0], Cmd, SubCmdPanel, verbRemove, "<buton-name>", "[<tab-name>]", "<panel-name>")
	usagePanelHelp     = fmt.Sprintf(usageCmdSubcmdVerbThisHelpF, os.Args[0], Cmd, SubCmdPanel, subCmdHelp)
)

func usageButton() (usage string) {
	lines := []string{
		UsageButtonAdd, usageButtonAddHelp, usageButtonRemove, usageButtonList, usageButtonHelp,
	}
	usage = newLine + strings.Join(lines, newLine)
	return
}

func usageTab() (usage string) {
	lines := []string{
		UsageTabAdd, usageTabAddHelp, usageTabRemove, usageTabHelp,
	}
	usage = newLine + strings.Join(lines, newLine)
	return
}

func usagePanel() (usage string) {
	lines := []string{
		UsagePanelAdd, usagePanelAddHelp, UsagePanelRemove, usagePanelHelp,
	}
	usage = newLine + strings.Join(lines, newLine)
	return
}

func usageCreateClear() (usage string) {
	lines := []string{
		UsageCreate, usageCreateHelp, usageClear,
	}
	usage = newLine + strings.Join(lines, newLine)
	return
}

func Usage() (usage string) {
	intro := `👀 THE FRONT END:
The frontend is made with a single button-pad, tab-bars and panel-groups.
* Button-Pad:
  Your app begins with the home panel which is the button-pad. The button labels give the user a general idea of what he or she can do with the app.
  When the user clicks a button, the view switches to the button's tab-bar panel or to the button's panel group.
* Tab-bars:
  In a tab-bar each tab has it's own panel group.
* Panel-groups:
  In a panel-group only one panel is visible at a time.
  Each panel in a panel-group has it's own unique design and purpose that you control.
  The framework gives each panel-group:
    * it's own messenger which communicates with the back-end.

MANAGING THE FRONT END:
`
	usage = intro +
		usageCreateClear() + newLine +
		usageButton() + newLine +
		usageTab() + newLine +
		usagePanel() + newLine
	return
}
