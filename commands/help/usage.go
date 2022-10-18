package help

import (
	"github.com/josephbudd/kickfyne/commands/framework"
	"github.com/josephbudd/kickfyne/commands/frontend"
	"github.com/josephbudd/kickfyne/commands/message"
	"github.com/josephbudd/kickfyne/commands/record"
	"github.com/josephbudd/kickfyne/commands/version"
)

const (
	usageCmdSubcmdF = "$ %s %s %s"
	newLine         = "\n"
	newParagraph    = "\n\n"
	gettingStarted  = `🍻 INTRODUCING kickfyne!
kickfyne is a tool to help build an application using the fyne toolkit which has among other things a very nice GUI. The fyne toolkit web site is located at https://fyne.io/. This kickfyne project is not in any way associated with this fyne toolkit project.
`
)

func Usage() (usage string) {
	usage =
		version.V() + newParagraph +
			gettingStarted + newParagraph +
			framework.Usage() + newParagraph +
			frontend.Usage() + newParagraph +
			message.Usage() + newParagraph +
			record.Usage() + newParagraph
	return
}
