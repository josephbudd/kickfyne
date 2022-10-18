package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ScreenHasAppTabsPanel(folderPaths *FolderPaths, screenPackageName string) (has bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	fPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName, AppTabsPanelFileName)
	if _, err = os.Stat(fPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	has = true
	return
}

func ScreenHasDocTabsPanel(folderPaths *FolderPaths, screenPackageName string) (has bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	fPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName, DocTabsPanelFileName)
	if _, err = os.Stat(fPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	has = true
	return
}

func ScreenHasAccordionPanel(folderPaths *FolderPaths, screenPackageName string) (has bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	fPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName, AccordionPanelFileName)
	if _, err = os.Stat(fPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	has = true
	return
}

// ScreenPackageNames returns the names of the screen packages.
func ScreenPackageNames(folderPaths *FolderPaths) (screenNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ScreenNames: %w", err)
		}
	}()

	var folderNames []string
	folderNames, err = FolderNames(folderPaths.FrontendGUIScreens)
	screenNames = make([]string, 0, len(folderNames))
	for _, fileName := range folderNames {
		if fileName == strings.ToLower(fileName) {
			screenNames = append(screenNames, fileName)
		}
	}
	return
}

// IsCurrentScreenName returns if the screenPackageName exists.
func IsCurrentScreenName(screenPackageName string, folderPaths *FolderPaths) (is bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.IsCurrentScreenName: %w", err)
		}
	}()

	var packageNames []string
	if packageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	for _, packageName := range packageNames {
		if is = packageName == screenPackageName; is {
			return
		}
	}
	return
}

// ValidateNewScreenPackageName validates a new panel name for a screenName.
func ValidateNewScreenPackageName(
	screenPackageName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateNewScreenPackageName: %w", err)
		}
	}()

	if screenPackageName != strings.ToLower(screenPackageName) {
		// The screen package name must be lower case.
		failureMessage = "The screen package name must be lower case."
		return
	}
	var currentPackageNames []string
	if currentPackageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	for _, currentPackageName := range currentPackageNames {
		if currentPackageName == screenPackageName {
			// Not a new screen package name.
			failureMessage = fmt.Sprintf("The screen package name %q is not a new screen package name.", screenPackageName)
			return
		}
	}
	// The package name is new and valid.
	isValid = true
	return
}

// ValidateCurrentScreenPackageName validates a new panel name for a screenName.
func ValidateCurrentScreenPackageName(
	screenPackageName string,
	allowLandingScreenName bool,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateCurrentScreenPackageName: %w", err)
		}
	}()

	if screenPackageName != strings.ToLower(screenPackageName) {
		// The screen package name must be lower case.
		failureMessage = "The screen package name must be lower case."
		return
	}
	var currentPackageNames []string
	if currentPackageNames, err = ScreenPackageNames(folderPaths); err != nil {
		return
	}
	if !allowLandingScreenName {
		var metaData FyneAppMetaData
		if metaData, err = ReadMetaData(folderPaths); err != nil {
			return
		}
		if screenPackageName == metaData.FrontEnd.Landing {
			failureMessage = fmt.Sprintf("The screen package name %q is not allowed because that package is the landing screen.", screenPackageName)
			return
		}
	}
	for _, currentPackageName := range currentPackageNames {
		if currentPackageName == screenPackageName {
			// This is a current screen package name.
			isValid = true
			return
		}
	}
	// The package name is not a current screen package name.
	failureMessage = fmt.Sprintf("The screen package name %q is not a current screen package name.", screenPackageName)
	return
}
