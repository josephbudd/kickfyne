package frontend

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/gui/screens"
	"github.com/josephbudd/kickfyne/source/utils"
)

// handlePanel passes control to the correct handlers.
func handlePanel(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handlePanel: %w", err)
		}
	}()

	if !isBuilt {
		fmt.Println("The app must be initailized before the front end panels can be added or removed.")
		return
	}
	if len(args) == 1 {
		fmt.Println(UsagePanel())
		return
	}
	// args[0] is "panel"
	// args[1] is the verb
	switch args[1] {
	case verbAddAccordion:
		// args[0] is "panel"
		// args[1] is "add-accordion"
		// args[2] is the <screen-package-name>
		err = handleAddAccordionPanel(pathWD, args[2], importPrefix)
	case verbAddAppTabs:
		// args[0] is "panel"
		// args[1] is "add-apptabs"
		// args[2] is the <screen-package-name>
		err = handleAddAppTabsPanel(pathWD, args[2], importPrefix)
	case verbAddDocTabs:
		// args[0] is "panel"
		// args[1] is "add-apptabs"
		// args[2] is the <screen-package-name>
		err = handleAddDocTabsPanel(pathWD, args[2], importPrefix)
	case verbList:
		// args[0] is "panel"
		// args[1] is "list"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsagePanel())
			return
		}
		err = handlePanelList(pathWD, args[2], importPrefix)
	case verbAdd:
		// args[0] is "panel"
		// args[1] is "add"
		// args[2] is the <screen-package-name>
		// args[3] is the <panel-name>
		if len(args) != 4 {
			fmt.Println(UsagePanel())
			return
		}
		err = handlePanelAdd(pathWD, args[2], args[3], importPrefix)
	case verbRemove:
		// args[0] is "panel"
		// args[1] is "remove"
		// args[2] is the <screen-package-name>
		// args[3] is the <panel-name>
		if len(args) != 4 {
			fmt.Println(UsagePanel())
			return
		}
		err = handlePanelRemove(pathWD, args[2], args[3], importPrefix)
	case subCmdHelp:
		fmt.Println(UsagePanel())
	default:
		fmt.Println(UsagePanel())
	}
	return
}

func handleAddAccordionPanel(pathWD, screenPackageName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleAddAccordionPanel: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, false, folderPaths); !isValid || err != nil {
		return
	}
	// Make sure not an Accordion screen.
	var hasFile bool
	if hasFile, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an accordion panel file.", screenPackageName)
		return
	}
	// Make sure not an AppTabs screen.
	if hasFile, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an AppTabs panel file.", screenPackageName)
		return
	}
	// Make sure not n DocTabs screen.
	if hasFile, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has a DocTabs panel file.", screenPackageName)
		return
	}
	// Add the panel file.
	if err = screens.BuildAccordionPanelFile(
		screenPackageName,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessagePanelAdd(screenPackageName, utils.AccordionPanelName)
	return
}

func handleAddAppTabsPanel(pathWD, screenPackageName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleAddAppTabsPanel: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, false, folderPaths); !isValid || err != nil {
		return
	}
	// Make sure not an Accordion screen.
	var hasFile bool
	if hasFile, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an accordion panel file.", screenPackageName)
		return
	}
	// Make sure not an AppTabs screen.
	if hasFile, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an AppTabs panel file.", screenPackageName)
		return
	}
	// Make sure not n DocTabs screen.
	if hasFile, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has a DocTabs panel file.", screenPackageName)
		return
	}
	// Add the panel file.
	if err = screens.BuildAppTabsPanelFile(
		screenPackageName,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessagePanelAdd(screenPackageName, utils.AppTabsPanelName)
	return
}

func handleAddDocTabsPanel(pathWD, screenPackageName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleAddAppTabsPanel: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, false, folderPaths); !isValid || err != nil {
		return
	}
	// Make sure not an Accordion screen.
	var hasFile bool
	if hasFile, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an accordion panel file.", screenPackageName)
		return
	}
	// Make sure not an AppTabs screen.
	if hasFile, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has an AppTabs panel file.", screenPackageName)
		return
	}
	// Make sure not na DocTabs screen.
	if hasFile, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if hasFile {
		failureMessage = fmt.Sprintf("The screen package %q already has a DocTabs panel file.", screenPackageName)
		return
	}
	// Add the panel file.
	if err = screens.BuildDocTabsPanelFile(
		screenPackageName,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessagePanelAdd(screenPackageName, utils.AppTabsPanelName)
	return
}

func handlePanelAdd(pathWD, screenPackageName, screenPanelName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handlePanelAdd: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, true, folderPaths); !isValid || err != nil {
		return
	}
	// Validate the panel name.
	if isValid, failureMessage, err = utils.ValidateNewScreenPanelName(screenPackageName, screenPanelName, folderPaths); !isValid || err != nil {
		return
	}

	// Make sure not an accordion screen.
	var hasAccordionPanel bool
	var hasAppTabsPanel bool
	var hasDocTabsPanel bool
	if hasAccordionPanel, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if !hasAccordionPanel {
		// Make sure not an AppTabs screen.
		if hasAppTabsPanel, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
			return
		}
		if !hasAppTabsPanel {
			// Make sure not an DocTabs screen.
			if hasDocTabsPanel, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
				return
			}
		}
	}

	// Add the panel file.
	if err = screens.BuildPanelFile(
		screenPackageName,
		screenPanelName,
		hasAccordionPanel,
		hasAppTabsPanel,
		hasDocTabsPanel,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessagePanelAdd(screenPackageName, screenPanelName)
	return
}

// handlePanelRemove handles the removal of a panel from a screen package.
func handlePanelRemove(pathWD, screenPackageName, screenPanelName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handlePanelRemove: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, false, folderPaths); !isValid || err != nil {
		return
	}
	// Validate the panel name.
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPanelName(screenPackageName, screenPanelName, folderPaths); !isValid || err != nil {
		return
	}

	// Make sure not an accordion screen.
	var hasAccordionPanel bool
	var hasAppTabsPanel bool
	var hasDocTabsPanel bool
	if hasAccordionPanel, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if !hasAccordionPanel {
		// Make sure not an AppTabs screen.
		if hasAppTabsPanel, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
			return
		}
		if !hasAppTabsPanel {
			// Make sure not an DocTabs screen.
			if hasDocTabsPanel, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
				return
			}
		}
	}
	if err = screens.RemovePanelFile(screenPackageName, screenPanelName, hasAccordionPanel, hasAppTabsPanel, hasDocTabsPanel, importPrefix, folderPaths); err != nil {
		return
	}
	successMessage = successMessagePanelRemove(screenPackageName, screenPanelName)
	return
}

// handlePanelList handles the listing of the panels in a screen package.
func handlePanelList(pathWD, screenPackageName, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handlePanelList: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateCurrentScreenPackageName(screenPackageName, true, folderPaths); !isValid || err != nil {
		return
	}
	screenFolderPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName)
	var panelNames []string
	if panelNames, err = utils.PanelNames(screenFolderPath); err != nil {
		return
	}
	// Display the list.
	fmt.Printf("List of %d panels in the %q screen package.\n", len(panelNames), screenPackageName)
	for i, panelName := range panelNames {
		fmt.Printf("%d. %s\n", i+1, panelName)
	}
	return
}

func successMessagePanelAdd(screenPackageName, panelName string) (successMessage string) {
	panelFileName := utils.PanelFileName(panelName)
	panelRelativeFilePath := utils.PanelFileRelativeFilePath(screenPackageName, panelName)
	screenRelativeFilePath := utils.ScreenFileRelativeFilePath(screenPackageName)
	successMessage = fmt.Sprintf("Added the panel named %q to the screen package %q.", panelFileName, screenPackageName) +
		"\n" +
		fmt.Sprintf("KICKFYNE TODO: The panel file at %s may need some editing.\nKICKFYNE TODO: The screen file at %s may need some editing.\n", panelRelativeFilePath, screenRelativeFilePath)
	return
}

func successMessagePanelRemove(screenPackageName, panelName string) (successMessage string) {
	screenRelativeFilePath := utils.ScreenFileRelativeFilePath(screenPackageName)
	successMessage = fmt.Sprintf("Removed the panel named %q from the screen package %q.", utils.AppTabsPanelFileName, screenPackageName) +
		"\n" +
		fmt.Sprintf("KICKFYNE TODO: The screen file at %s may need some editing.\n", screenRelativeFilePath)
	return
}
