package frontend

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickfyne/loader/buttonsyaml"
	"github.com/josephbudd/kickfyne/source/utils"
)

func handleCreate(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleCreate: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			dumperCh <- "Failure:"
			dumperCh <- failureMessage
			dumperCh <- Usage()
		case len(successMessage) > 0:
			dumperCh <- "Success:"
			dumperCh <- successMessage
		}
	}()
	if len(args) == 1 {
		failureMessage = `Missing a file path or "help".`
		return
	}
	switch args[1] {
	case subCmdHelp:
		dumperCh <- createHelp
	default:
		if !isBuilt {
			failureMessage = utils.NoFrameworkMessage("the front end panels can be created.")
			return
		}
		// Read the buttons YAML file.
		var buttonsYAML buttonsyaml.YAML
		var isOK bool
		if buttonsYAML, isOK, failureMessage, err = buttonsyaml.Load(args[1], folderPaths); err != nil || !isOK {
			return
		}
		// Add each button.
		for buttonIndex, buttonYAMLButton := range buttonsYAML.Buttons {
			if err = addButton(buttonYAMLButton, buttonIndex, importPrefix, folderPaths); err != nil {
				return
			}
		}
		// Success message.
		lButtons := len(buttonsYAML.Buttons)
		var messages = make([]string, lButtons+1)
		messages[0] = fmt.Sprintf("All %d home buttons are created.", len(buttonsYAML.Buttons))
		for i, button := range buttonsYAML.Buttons {
			j := i + 1
			if len(button.Tabs) > 0 {
				messages[j] = fmt.Sprintf("%d. Button named %q and it's tab-bar and their panel-groups.", j, button.Name)
			} else {
				messages[j] = fmt.Sprintf("%d. Button named %q and it's panel-groups.", j, button.Name)
			}
		}
		successMessage = strings.Join(messages, "\n")
	}
	return
}
