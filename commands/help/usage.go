package help

import (
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
kickfyne is a cli that helps you build an application using the fyne toolkit which has among other things a very nice GUI.
The fyne tookkit web site is located at https://fyne.io/.
The kickfyne project is not in any way associated with the fyne toolkit project.

🗽 GETTING STARTED WITH kickfyne.
cd to the app folder: "$ cd ~/projects/<name of my app>"
run go mod init:      "$ go mod init "example.com/<name of my app>" or "github.com/me/<name of my app>"
create the framework: "$ kickfyne framework"

🚧 THE FRAMEWORK:
The framework is contained in 3 folders.
1. backend/ which contains the back end code.
2. frontend/ which contains the GUI code.
3. shared/ which contains shared code.

🔨 BUILDING THE APP:
You can build the app after running the command "$ kickfyne framework".
Note that main.go is in the shared/ folder.
The following build example is done in the app's folder.
$ go mod tidy
$ go build -o <name of executable> ./shared
$ ./<name of executable>
`
)

func Usage() (usage string) {
	usage =
		version.V() + newParagraph +
			gettingStarted + newParagraph +
			frontend.Usage() + newParagraph +
			message.Usage() + newParagraph +
			record.Usage() + newParagraph
	return
}
