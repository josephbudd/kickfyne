package frontend

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/gui/screens"
	"github.com/josephbudd/kickfyne/source/frontend/landingscreen"
	"github.com/josephbudd/kickfyne/source/root"
	"github.com/josephbudd/kickfyne/source/utils"
)

var errNil = fmt.Errorf("errNil")

// handleScreen passes control to the correct handlers.
func handleScreen(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	defer func() {
		if err == errNil {
			err = nil
			return
		}
		if err != nil {
			err = fmt.Errorf("frontend.handlePanel: %w", err)
		}
	}()

	if !isBuilt {
		fmt.Println("The application must be initialized before the front end panels can be added or removed.")
		return
	}
	if len(args) == 1 {
		fmt.Println(UsageScreen())
		return
	}

	// Need the application name when rebuilding main.go.
	appName := filepath.Base(importPrefix)

	// args[0] is "screen"
	// args[1] is the verb
	switch args[1] {
	case subCmdLanding:
		// args[0] is "screen"
		// args[1] is "landing"
		err = handleScreenLanding(pathWD, importPrefix)
	case verbList:
		// args[0] is "screen"
		// args[1] is "list"
		if len(args) != 2 {
			fmt.Println(UsageScreen())
			return
		}
		err = handleScreenList(pathWD, importPrefix)
	case verbAdd:
		// args[0] is "screen"
		// args[1] is "add"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is a package.\nKICKFYNE TODO: Correct this package doc commment.", args[2])
		if err = handleScreenAdd(pathWD, args[2], screenDocComment, importPrefix, appName); err != nil {
			return
		}
		fmt.Println(successMessageScreenAdd(
			args[2],
			false,
			false,
			false,
		))
	case verbAddAccordion:
		// args[0] is "screen"
		// args[1] is "add-accordion"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is an Accordion package.\nKICKFYNE TODO: Correct this package doc commment.", args[2])
		if err = handleScreenAdd(pathWD, args[2], screenDocComment, importPrefix, appName); err != nil {
			return
		}
		var folderPaths *utils.FolderPaths
		if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// Add the panel file.
		if err = screens.BuildAccordionPanelFile(
			args[2],
			importPrefix,
			folderPaths,
		); err != nil {
			return
		}
		fmt.Println(successMessageScreenAdd(
			args[2],
			true,
			false,
			false,
		))
	case verbAddAppTabs:
		// args[0] is "screen"
		// args[1] is "add-accordion"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is an AppTabs package.\nKICKFYNE TODO: Correct this package doc commment.", args[2])
		if err = handleScreenAdd(pathWD, args[2], screenDocComment, importPrefix, appName); err != nil {
			return
		}
		var folderPaths *utils.FolderPaths
		if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// Add the panel file.
		if err = screens.BuildAppTabsPanelFile(
			args[2],
			importPrefix,
			folderPaths,
		); err != nil {
			return
		}
		fmt.Println(successMessageScreenAdd(
			args[2],
			false,
			true,
			false,
		))
	case verbAddDocTabs:
		// args[0] is "screen"
		// args[1] is "add-accordion"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		screenDocComment := fmt.Sprintf("Package %s is an DocTabs package.\nKICKFYNE TODO: Correct this package doc commment.", args[2])
		if err = handleScreenAdd(pathWD, args[2], screenDocComment, importPrefix, appName); err != nil {
			return
		}
		var folderPaths *utils.FolderPaths
		if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// Add the panel file.
		if err = screens.BuildDocTabsPanelFile(
			args[2],
			importPrefix,
			folderPaths,
		); err != nil {
			return
		}
		fmt.Println(successMessageScreenAdd(
			args[2],
			false,
			false,
			true,
		))
	case verbRemove:
		// args[0] is "screen"
		// args[1] is "remove"
		// args[2] is the <screen-package-name>
		if len(args) != 3 {
			fmt.Println(UsageScreen())
			return
		}
		err = handleScreenRemove(pathWD, args[2], importPrefix, appName)
	case subCmdHelp:
		// args[0] is "screen"
		// args[1] is "help"
		fmt.Println(UsageScreen())
	default:
		// args[0] is "screen"
		fmt.Println(UsageScreen())
	}
	return
}

// handleScreenLanding handles rebuilding frontend/screenlanding/landing.go.
func handleScreenLanding(pathWD, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenLanding: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
			err = errNil
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}

	// Get the application's meta data.
	var metaData utils.FyneAppMetaData
	if metaData, err = utils.ReadMetaData(folderPaths); err != nil {
		return
	}
	// Confirm that the meta data file uses the name of an existing screen package.
	var isValid bool
	if isValid, err = utils.IsCurrentScreenName(metaData.FrontEnd.Landing, folderPaths); err != nil {
		return
	}
	if !isValid {
		// Param metaData.FrontEnd.Landing is not the name of an existing screen.
		failureMessage = fmt.Sprintf("The [Fronttend].Landing from ./%s is not a valid.\n%q is not a screen package name.", utils.FyneAppTOMLFileName, metaData.FrontEnd.Landing)
		return
	}
	// Rebuild ./frontend/landingscreen/landing.go
	if err = landingscreen.BuildLanding(importPrefix, folderPaths); err != nil {
		return
	}
	successMessage = fmt.Sprintf("The screen package %q is now the landing screen.", metaData.FrontEnd.Landing)
	return
}

// handleScreenAdd handles adding a screen package.
func handleScreenAdd(pathWD, screenPackageName, screenPackageDoc, importPrefix, appName string) (err error) {

	var failureMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenAdd: %w", err)
			return
		}
		if len(failureMessage) > 0 {
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
			err = errNil
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Validate the screen package name.
	var isValid bool
	if isValid, failureMessage, err = utils.ValidateNewScreenPackageName(screenPackageName, folderPaths); !isValid || err != nil {
		return
	}
	// Create the package folder with no panels.
	if err = screens.BuildPackageWithoutPanels(
		screenPackageName,
		screenPackageDoc,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}
	// Rebuild main.go.
	var screenPackageNames []string
	if screenPackageNames, err = utils.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	root.RebuildMainScreensGo(appName, screenPackageNames, importPrefix, folderPaths)
	return
}

// handleScreenRemove handles the removal of a screen package.
func handleScreenRemove(pathWD, screenPackageName, importPrefix, appName string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenRemove: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
			err = errNil
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
	var msg string
	if isValid, msg, err = utils.ValidateCurrentScreenPackageName(screenPackageName, false, folderPaths); !isValid || err != nil {
		if !isValid {
			failureMessage = msg
		}
		return
	}
	packageFolderPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName)
	if _, err = os.Stat(packageFolderPath); err != nil {
		if os.IsNotExist(err) {
			// The folder does not exist.
			err = nil
			successMessage = fmt.Sprintf("The screen package %q was previously removed for some reason.", screenPackageName)
		}
		return
	}
	if err = os.RemoveAll(packageFolderPath); err != nil {
		return
	}
	// Removed the folder.
	successMessage = fmt.Sprintf("Removed the screen package %q.", screenPackageName)
	// Rebuild main.go.
	var screenPackageNames []string
	if screenPackageNames, err = utils.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	root.RebuildMainScreensGo(appName, screenPackageNames, importPrefix, folderPaths)
	return
}

// handleScreenList handles the listing of the screen packages.
func handleScreenList(pathWD, importPrefix string) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleScreenList: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			fmt.Println("Failure:")
			fmt.Println(failureMessage)
			err = errNil
		case len(successMessage) > 0:
			fmt.Println("Success:")
			fmt.Println(successMessage)
		}
	}()

	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	// Get the screen names.
	var screenNames []string
	if screenNames, err = utils.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	// Get the panel names.
	screenPanelNames := make(map[string][]string)
	for _, screenName := range screenNames {
		screenFolderPath := filepath.Join(folderPaths.FrontendGUIScreens, screenName)
		if screenPanelNames[screenName], err = utils.PanelNames(screenFolderPath); err != nil {
			return
		}
	}
	// Display the list.
	fmt.Printf("List of %d screen packages.\n", len(screenNames))
	for i, screenName := range screenNames {
		panelNames := screenPanelNames[screenName]
		lPanels := len(panelNames)
		if lPanels == 0 {
			fmt.Printf("% d. The %s screen has no panels.\n", i+1, screenName)
			continue
		}
		// Display the panels
		fmt.Printf("% d. %s: %s.\n", i+1, screenName, utils.ScreenFileRelativeFilePath(screenName))
		// fmt.Printf("% d. %s: %sThe %s screen has %d panels.\n", i+1, screenName, lPanels)
		for j, panelName := range panelNames {
			fmt.Printf("    % d. %s: %s.\n", j+1, panelName, utils.PanelFileRelativeFilePath(screenName, panelName))
		}
	}

	return
}

func successMessageScreenAdd(
	screenPackageName string,
	isAccordionScreen bool,
	isAppTabScreen bool,
	isDocTabScreen bool,
) (successMessage string) {
	var defaultPanelType string
	switch {
	case isAccordionScreen:
		defaultPanelType = "Accordion"
	case isAppTabScreen:
		defaultPanelType = "AppTabs"
	case isDocTabScreen:
		defaultPanelType = "DocTabs"
	default:
		defaultPanelType = "generic"
	}
	docRelativeFilePath := utils.DocFileRelativeFilePath(screenPackageName)
	successMessage = fmt.Sprintf("Added the %s screen package named %q.", defaultPanelType, screenPackageName) +
		fmt.Sprintf("\nKICKFYNE TODO: The doc file at %s may need some editing.", docRelativeFilePath) +
		defaultPanelReviewMessage(screenPackageName, isAccordionScreen, isAppTabScreen, isDocTabScreen)
	return
}
