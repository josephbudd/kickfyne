package frontend

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/loader/buttonyaml"
	"github.com/josephbudd/kickfyne/source/utils"
)

func handleButton(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleButton: %w", err)
		}
	}()

	if !isBuilt {
		dumperCh <- "The app must be initailized before the front end panels can be removed."
		return
	}
	if len(args) == 1 {
		dumperCh <- "Missing a verb."
		dumperCh <- usageButton()
		return
	}
	args = args[1:]
	switch args[0] {
	case verbAdd:
		// button add <path to button yaml file>
		if len(args) == 1 {
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- usageButton()
			return
		}
		switch args[1] {
		case subCmdHelp:
			dumperCh <- buttonAddHelp
			return
		case "":
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- usageButton()
		default:
			err = handleButtonAdd(args[1], importPrefix, dumperCh, folderPaths)
		}
	case verbRemove:
		// button remove <button name>
		if len(args) == 1 {
			dumperCh <- `Missing a button name.`
			dumperCh <- usageButton()
			return
		}
		err = handleButtonRemove(args[1], importPrefix, dumperCh, folderPaths)
	case subCmdHelp:
		dumperCh <- usageButton()
	}
	return
}

func handleButtonRemove(buttonName string, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleButtonRemove: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			dumperCh <- "Failure:"
			dumperCh <- failureMessage
		case len(successMessage) > 0:
			dumperCh <- "Success:"
			dumperCh <- successMessage
		}
	}()

	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentButtonName(buttonName, folderPaths); !isValid || err != nil {
		return
	}
	var buttonHasTabBar bool
	if buttonHasTabBar, err = utils.ButtonHasTabBar(buttonName, folderPaths); err != nil {
		return
	}
	var filePath string
	var folderPath string
	if filePath, folderPath, err = utils.ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	if err = os.Remove(filePath); err != nil {
		return
	}
	if err = os.RemoveAll(folderPath); err != nil {
		return
	}
	if buttonHasTabBar {
		successMessage = fmt.Sprintf("Button named %q and it's tab-bars and panel-groups have been removed.", buttonName)
	} else {
		successMessage = fmt.Sprintf("Button named %q and panel-groups have been removed.", buttonName)
	}
	return
}

func handleButtonAdd(path string, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleButtonAdd: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			dumperCh <- "Failure:"
			dumperCh <- failureMessage
		case len(successMessage) > 0:
			dumperCh <- "Success:"
			dumperCh <- successMessage
		}
	}()

	// Read the button YAML file.
	var buttonYAML buttonyaml.YAML
	var isOK bool
	if buttonYAML, isOK, failureMessage, err = buttonyaml.Load(path, folderPaths); err != nil || !isOK {
		return
	}
	// Add the button.
	if err = addButton(buttonYAML.Button, buttonYAML.InsertBeforeButton, importPrefix, folderPaths); err != nil {
		return
	}
	// Build the success message.
	if len(buttonYAML.Button.Tabs) > 0 {
		successMessage = fmt.Sprintf("Button named %q and it's tab-bars and panel-groups have been added.", buttonYAML.Button.Name)
	} else {
		successMessage = fmt.Sprintf("Button named %q and panel-groups have been added.", buttonYAML.Button.Name)
	}
	return
}
