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
		fmt.Println("The application must be initialized before the front end panels can be added or removed.")
		return
	}
	if len(args) == 1 {
		fmt.Println(UsagePanel())
		return
	}
	// args[0] is "panel"
	// args[1] is the verb
	switch args[1] {
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

	// Determine the kind of screen.
	var isAccordionScreen bool
	var isAppTabScreen bool
	var isDocTabScreen bool
	if isAccordionScreen, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if !isAccordionScreen {
		// Make sure not an AppTabs screen.
		if isAppTabScreen, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
			return
		}
		if !isAppTabScreen {
			// Make sure not an DocTabs screen.
			if isDocTabScreen, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
				return
			}
		}
	}

	// Add the panel file.
	if err = screens.BuildPanelFile(
		screenPackageName,
		screenPanelName,
		isAccordionScreen,
		isAppTabScreen,
		isDocTabScreen,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	successMessage = successMessagePanelAdd(screenPackageName, screenPanelName, isAccordionScreen, isAppTabScreen, isDocTabScreen)
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

	// Get the screen type.
	var isAccordionScreen bool
	var isAppTabScreen bool
	var isDocTabScreen bool
	if isAccordionScreen, err = utils.ScreenHasAccordionPanel(folderPaths, screenPackageName); err != nil {
		return
	}
	if !isAccordionScreen {
		// Make sure not an AppTabs screen.
		if isAppTabScreen, err = utils.ScreenHasAppTabsPanel(folderPaths, screenPackageName); err != nil {
			return
		}
		if !isAppTabScreen {
			// Make sure not an DocTabs screen.
			if isDocTabScreen, err = utils.ScreenHasDocTabsPanel(folderPaths, screenPackageName); err != nil {
				return
			}
		}
	}
	if err = screens.RemovePanelFile(screenPackageName, screenPanelName, isAccordionScreen, isAppTabScreen, isDocTabScreen, importPrefix, folderPaths); err != nil {
		return
	}
	successMessage = successMessagePanelRemove(screenPackageName, screenPanelName, isAccordionScreen, isAppTabScreen, isDocTabScreen)
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
		fmt.Printf("    % d. %s: %s.\n", i+1, panelName, utils.PanelFileRelativeFilePath(screenPackageName, panelName))
	}
	return
}

func successMessagePanelAdd(screenPackageName, panelName string, isAccordionScreen, isAppTabScreen, isDocTabScreen bool) (successMessage string) {
	panelFileRelativePath := utils.PanelFileRelativeFilePath(screenPackageName, panelName)
	successMessage = fmt.Sprintf("Added the panel named %q to the screen package %q.", panelName, screenPackageName) +
		fmt.Sprintf("\nKICKFYNE TODO: Edit the new panel file at %s.", panelFileRelativePath) +
		defaultPanelEditingMessage(screenPackageName, isAccordionScreen, isAppTabScreen, isDocTabScreen)
	return
}

func successMessagePanelRemove(screenPackageName, panelName string, isAccordionScreen, isAppTabScreen, isDocTabScreen bool) (successMessage string) {
	successMessage = fmt.Sprintf("Removed the panel named %q from the screen package %q.", panelName, screenPackageName) +
		defaultPanelEditingMessage(screenPackageName, isAccordionScreen, isAppTabScreen, isDocTabScreen)
	return
}

func defaultPanelEditingMessage(
	screenPackageName string,
	isAccordionScreen bool,
	isAppTabScreen bool,
	isDocTabScreen bool,
) (successMessage string) {
	var defaultPanelName string
	var defaultPanelType string
	switch {
	case isAccordionScreen:
		defaultPanelName = utils.AccordionPanelName
		defaultPanelType = "Accordion"
	case isAppTabScreen:
		defaultPanelName = utils.AppTabsPanelName
		defaultPanelType = "AppTabs"
	case isDocTabScreen:
		defaultPanelName = utils.DocTabsPanelName
		defaultPanelType = "DocTabs"
	default:
		return
	}
	defaultPanelFileRelativePath := utils.PanelFileRelativeFilePath(screenPackageName, defaultPanelName)
	successMessage = "\n" + fmt.Sprintf("KICKFYNE TODO: The %s panel file at %s may need some editing.", defaultPanelType, defaultPanelFileRelativePath)
	return
}

func defaultPanelReviewMessage(
	screenPackageName string,
	isAccordionScreen bool,
	isAppTabScreen bool,
	isDocTabScreen bool,
) (successMessage string) {
	var defaultPanelName string
	var defaultPanelType string
	var funcName string
	switch {
	case isAccordionScreen:
		defaultPanelName = utils.AccordionPanelName
		defaultPanelType = "Accordion"
		funcName = "accordionItems"
	case isAppTabScreen:
		defaultPanelName = utils.AppTabsPanelName
		defaultPanelType = "AppTabs"
		funcName = "tabItems"
	case isDocTabScreen:
		defaultPanelName = utils.DocTabsPanelName
		defaultPanelType = "DocTabs"
		funcName = "tabItems"
	default:
		return
	}
	defaultPanelFileRelativePath := utils.PanelFileRelativeFilePath(screenPackageName, defaultPanelName)
	successMessage = "\n" + fmt.Sprintf("KICKFYNE TODO: You may want to review func %s in the %s panel file at %s.", funcName, defaultPanelType, defaultPanelFileRelativePath)
	return
}
