package frontend

import (
	"github.com/josephbudd/kickfyne/loader/buttonsyaml"
	"github.com/josephbudd/kickfyne/loader/buttonyaml"
	"github.com/josephbudd/kickfyne/loader/panelyaml"
	"github.com/josephbudd/kickfyne/loader/tabyaml"
)

const (
	helpBlankLine = "\n\n"

	gear   = "⚙"
	laptop = "💻"
)

var (
	panelAddHelp  string
	tabAddHelp    string
	buttonAddHelp string
	createHelp    string
)

func init() {
	panelAddHelp = "\nADD A PANEL TO A BUTTON OR TAB.\n\n" + laptop + " COMMAND LINE:\n" + UsagePanelAdd + helpBlankLine + panelyaml.AddPanelYAMLExample
	tabAddHelp = "\nADD A TAB TO A BUTTON.\n\n" + laptop + " COMMAND LINE:\n" + UsageTabAdd + helpBlankLine + tabyaml.AddTabYAMLExample
	buttonAddHelp = "\nADD A BUTTON.\n\n" + laptop + " COMMAND LINE:\n" + UsageButtonAdd + helpBlankLine + buttonyaml.AddButtonYAMLExamples
	createHelp = "\nCREATE BUTTONS, TABS AND PANELS.\n\n" + laptop + " COMMAND LINE:\n" + UsageCreate + helpBlankLine + buttonsyaml.CreateButtonsYAMLExample
}
