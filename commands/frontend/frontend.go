package frontend

import (
	"fmt"
)

const (
	CmdPanel         = "panel"
	CmdScreen        = "screen"
	subCmdHelp       = "help"
	subCmdLanding    = "update-landing"
	verbAdd          = "add"
	verbRemove       = "remove"
	verbList         = "list"
	verbAddAppTabs   = "add-apptabs"
	verbAddDocTabs   = "add-doctabs"
	verbAddAccordion = "add-accordion"
)

// Handler passes control to the correct handlers.
func Handler(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	if !isBuilt || len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Handler: %w", err)
		}
	}()

	switch args[0] {
	case CmdScreen:
		err = handleScreen(pathWD, args, isBuilt, importPrefix)
	case CmdPanel:
		err = handlePanel(pathWD, args, isBuilt, importPrefix)
	case subCmdHelp:
		fmt.Println(Usage())
	default:
		fmt.Println(Usage())
	}
	return
}
