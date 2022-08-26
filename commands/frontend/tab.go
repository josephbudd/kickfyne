package frontend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/loader/tabyaml"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar/tabpanelgroup"
	"github.com/josephbudd/kickfyne/source/utils"
)

func handleTab(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleTab: %w", err)
		}
	}()

	if !isBuilt {
		dumperCh <- "The app must be initailized before the front end tabs can be added or removed."
		return
	}
	if len(args) == 1 {
		dumperCh <- "Missing a verb."
		dumperCh <- usageTab()
		return
	}
	args = args[1:]
	switch args[0] {
	case verbAdd:
		// tab add <path to button yaml file>
		if len(args) == 1 {
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- usageTab()
			return
		}
		switch args[1] {
		case subCmdHelp:
			dumperCh <- tabAddHelp
			return
		case "":
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- usageTab()
			return
		default:
			err = handleTabAdd(args[1], importPrefix, dumperCh, folderPaths)
		}
	case verbRemove:
		// tab remove <tab name>
		if len(args) == 1 {
			dumperCh <- "Missing a button name and a tab name."
			dumperCh <- usageTab()
			return
		}
		if len(args) == 2 {
			dumperCh <- "Missing a tab name."
			dumperCh <- usageTab()
			return
		}
		err = handleTabRemove(args[1], args[2], importPrefix, dumperCh, folderPaths)
	case subCmdHelp:
		dumperCh <- usageTab()
	}
	return
}

func handleTabRemove(buttonName, tabName, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleTabRemove: %w", err)
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
	var msg string
	if isValid, msg, err = utils.ValidateCurrentButtonName(buttonName, folderPaths); !isValid || err != nil {
		failureMessage = msg
		return
	}
	if isValid, msg, err = utils.ValidateCurrentTabName(buttonName, tabName, folderPaths); !isValid || err != nil {
		failureMessage = msg
		return
	}
	var folderPath string
	if _, folderPath, err = utils.ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	fileName := utils.TabFileName(tabName)
	tabFilePath := filepath.Join(folderPath, fileName)
	fileExists := true
	if err = os.Remove(tabFilePath); err != nil {
		if os.IsNotExist(err) {
			fileExists = false
			err = nil
		}
		if err != nil {
			return
		}
	}
	// The panel group folder.
	folderName := utils.TabPanelGroupFolderName(tabName)
	panelGroupFolderPath := filepath.Join(folderPath, folderName)
	folderExists := true
	if _, err = os.Stat(panelGroupFolderPath); err != nil {
		if os.IsNotExist(err) {
			folderExists = false
			err = nil
		}
		if err != nil {
			return
		}
	}
	if folderExists {
		if err = os.RemoveAll(panelGroupFolderPath); err != nil {
			return
		}
	}
	// Build the success message.
	var fileMessage string
	if fileExists {
		fileMessage = fmt.Sprintf("1. Removed the tab file at\n   %q.", tabFilePath)
	} else {
		fileMessage = fmt.Sprintf("1. The tab file at %q was previously removed for some reason.", tabFilePath)
	}
	var folderMessage string
	if folderExists {
		folderMessage = fmt.Sprintf("2. Removed the tab's panel group folder at\n   %q.", panelGroupFolderPath)
	} else {
		folderMessage = fmt.Sprintf("2. The tab's panel group folder at %q was previously removed for some reason.", panelGroupFolderPath)
	}
	msgs := []string{
		fmt.Sprintf("Removed the tab named %q from the button named %q.", tabName, buttonName),
		fileMessage,
		folderMessage,
	}
	successMessage = strings.Join(msgs, "\n")
	return
}

func handleTabAdd(path string, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleTabAdd: %w", err)
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
	var tabYAML tabyaml.YAML
	var isOK bool
	var msg string
	if tabYAML, isOK, msg, err = tabyaml.Load(path, folderPaths); err != nil || !isOK {
		if isOK {
			successMessage = msg
		} else {
			failureMessage = msg
		}
		return
	}
	if err = addTab(tabYAML.ButtonName, tabYAML.Tab, tabYAML.InsertBeforeTab, importPrefix, folderPaths); err != nil {
		return
	}
	return
}

func addTab(buttonName string, tab loader.Tab, tabIndex int, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.buttonName: %w", err)
		}
	}()

	var tabGroupName string
	var tabFolderPath string
	if tabGroupName, tabFolderPath, err = buttontabbar.BuildTab(
		utils.HomeGroupName,
		buttonName,
		tab.Name,
		tab.Label,
		tabIndex,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	defaultPanel := tab.Panels[0]
	if err = tabpanelgroup.Build(
		buttonName,
		tabGroupName,
		tab.Name,
		defaultPanel.Name,
		tabFolderPath,
		importPrefix,
	); err != nil {
		return
	}
	for _, panel := range tab.Panels {
		if err = tabpanelgroup.BuildPanel(
			buttonName,
			tabGroupName,
			tab.Name,
			tabFolderPath,
			panel.Name,
			panel.Description,
			panel.Heading,
			importPrefix,
		); err != nil {
			return
		}
	}
	return
}
