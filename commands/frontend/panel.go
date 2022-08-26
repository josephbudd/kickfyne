package frontend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/loader/panelyaml"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttonpanelgroup"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar/tabpanelgroup"
	"github.com/josephbudd/kickfyne/source/utils"
)

func handlePanel(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handlePanel: %w", err)
		}
	}()

	if !isBuilt {
		dumperCh <- "The app must be initailized before the front end panels can be added or removed."
		return
	}
	// panel add <file path>
	// panel remove <b> [<t>] <p>
	if len(args) == 1 {
		dumperCh <- "Missing a verb."
		dumperCh <- Usage()
		return
	}
	// args[0] is "panel"
	// args[1] is the verb
	switch args[1] {
	case verbAdd:
		// panel add <path to button yaml file>
		if len(args) == 2 {
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- Usage()
			return
		}
		// args[0] is "panel"
		// args[1] is the verb
		// args[2] is the path
		switch args[2] {
		case subCmdHelp:
			dumperCh <- panelAddHelp
			return
		case "":
			dumperCh <- `Missing a file path or "help".`
			dumperCh <- Usage()
		default:
			err = handlePanelAdd(args[2], importPrefix, dumperCh, folderPaths)
		}
	case verbRemove:
		// panel remove <button name> <tab name> <panel name>
		// panel remove <button name> <panel name>
		switch len(args) {
		case 2:
			// args[0] is "panel"
			// args[1] is the verb
			dumperCh <- "Missing a button name and a tab name."
			dumperCh <- Usage()
		case 3:
			// args[0] is "panel"
			// args[1] is the verb
			// args[2] is button name.
			dumperCh <- "Missing a tab name and or a panel name."
			dumperCh <- Usage()
		case 4:
			// args[0] is "panel"
			// args[1] is the verb
			// args[2] is button name.
			// args[3] is the panel name.
			err = handleButtonPanelRemove(args[2], args[3], importPrefix, dumperCh, folderPaths)
		case 5:
			// args[0] is "panel"
			// args[1] is the verb
			// args[2] is button name.
			// args[3] is the tab name.
			// args[4] is the panel name.
			err = handleTabPanelRemove(args[2], args[3], args[4], importPrefix, dumperCh, folderPaths)
		}
	case subCmdHelp:
		dumperCh <- usagePanel()
	}
	return
}

func handleTabPanelRemove(buttonName, tabName, panelName, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

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
		if !isValid {
			failureMessage = msg
		}
		return
	}
	if isValid, msg, err = utils.ValidateCurrentTabName(buttonName, tabName, folderPaths); !isValid || err != nil {
		if !isValid {
			failureMessage = msg
		}
		return
	}
	if isValid, msg, err = utils.ValidateCurrentTabPanelName(buttonName, tabName, panelName, folderPaths); !isValid || err != nil {
		if !isValid {
			failureMessage = msg
		}
		return
	}
	var tabBarFolderPath string
	if _, tabBarFolderPath, err = utils.ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	panelFilePath := filepath.Join(
		tabBarFolderPath,
		utils.TabPanelGroupFolderName(tabName),
		utils.PanelFileName(panelName),
	)

	fileExists := true
	if err = os.Remove(panelFilePath); err != nil {
		if os.IsNotExist(err) {
			fileExists = false
			err = nil
		}
		if err != nil {
			return
		}
	}
	// Build the success message.
	if fileExists {
		successMessage = fmt.Sprintf("Removed the panel named %q, from to the tab named %q, belonging to the button named %q.", panelName, tabName, buttonName)
	} else {
		successMessage = fmt.Sprintf("The panel named %q, belonging to the tab named %q, belonging to the button named %q was previously removed for some reason.", panelName, tabName, buttonName)
	}
	return
}

func handleButtonPanelRemove(buttonName, panelName, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

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
		if !isValid {
			failureMessage = msg
		}
		return
	}
	if isValid, msg, err = utils.ValidateCurrentButtonPanelName(buttonName, panelName, folderPaths); !isValid || err != nil {
		if !isValid {
			failureMessage = msg
		}
		return
	}
	// The panel group folder.
	var panelGroupFolderPath string
	if _, panelGroupFolderPath, err = utils.ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	// The panel file.
	fileName := utils.PanelFileName(panelName)
	panelFilePath := filepath.Join(panelGroupFolderPath, fileName)
	fileExists := true
	if err = os.Remove(panelFilePath); err != nil {
		if os.IsNotExist(err) {
			fileExists = false
			err = nil
		}
		if err != nil {
			return
		}
	}

	// Build the success message.
	if fileExists {
		successMessage = fmt.Sprintf("Removed the panel named %q, from to the button named %q.", panelName, buttonName)
	} else {
		successMessage = fmt.Sprintf("The panel named %q, belonging to the button named %q was previously removed for some reason.", panelName, buttonName)
	}
	return
}

func handlePanelAdd(path string, importPrefix string, dumperCh chan string, folderPaths *utils.FolderPaths) (err error) {

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
	var panelYAML panelyaml.YAML
	var isOK bool
	var msg string
	if panelYAML, isOK, msg, err = panelyaml.Load(path, folderPaths); err != nil || !isOK {
		if !isOK {
			failureMessage = msg
		}
		return
	}

	// Build the panel group folder path.
	var panelGroupFolderPath string
	if _, panelGroupFolderPath, err = utils.ButtonFileFolderPaths(panelYAML.ButtonName, folderPaths); err != nil {
		return
	}
	var isTabBar = len(panelYAML.TabName) > 0
	if isTabBar {
		tabPanelGroupFolderName := utils.TabPanelGroupFolderName(panelYAML.TabName)
		panelGroupFolderPath = filepath.Join(panelGroupFolderPath, tabPanelGroupFolderName)

	}
	var currentPanelNames []string
	if currentPanelNames, err = utils.PanelNames(panelGroupFolderPath); err != nil {
		return
	}
	msgs := make([]string, 0, len(currentPanelNames))
	for _, currentPanelName := range currentPanelNames {
		if currentPanelName == panelYAML.Panel.Name {
			if isTabBar {
				msg = fmt.Sprintf("The button named %q, has a tab named %q, which already has a panel named %q.", panelYAML.ButtonName, panelYAML.TabName, panelYAML.Panel.Name)
				msgs = append(msgs, msg)
			} else {
				msg = fmt.Sprintf("The button named %q, already has a panel named %q.", panelYAML.ButtonName, panelYAML.Panel.Name)
				msgs = append(msgs, msg)
			}
		}
	}
	if len(msgs) > 0 {
		failureMessage = strings.Join(msgs, "\n")
	}
	if !isTabBar {
		if err = addButtonPanel(panelYAML.ButtonName, panelYAML.Panel, importPrefix, folderPaths); err != nil {
			return
		}
		successMessage = fmt.Sprintf("Added the panel named %q the button named %q.", panelYAML.Panel.Name, panelYAML.ButtonName)
	} else {
		if err = addTabPanel(panelYAML.ButtonName, panelYAML.TabName, panelYAML.Panel, importPrefix, folderPaths); err != nil {
			return
		}
		successMessage = fmt.Sprintf("Added the panel named %q to the tab named %q belonging to the button named %q's.", panelYAML.Panel.Name, panelYAML.TabName, panelYAML.ButtonName)
	}
	return
}

func addButtonPanel(buttonName string, panel loader.Panel, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.buttonName: %w", err)
		}
	}()

	err = buttonpanelgroup.BuildPanel(
		panel.Name,
		panel.Description,
		panel.Heading,
		utils.HomeGroupName,
		buttonName,
		importPrefix,
		folderPaths,
	)
	return
}

func addTabPanel(buttonName, tabName string, panel loader.Panel, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.addTabPanel: %w", err)
		}
	}()

	_, tabBarPGroupName, tabBarFolderPath := buttontabbar.Names(buttonName, utils.HomeGroupName, folderPaths)
	err = tabpanelgroup.BuildPanel(
		buttonName,
		tabBarPGroupName,
		tabName,
		tabBarFolderPath,
		panel.Name,
		panel.Description,
		panel.Heading,
		importPrefix,
	)
	return
}
