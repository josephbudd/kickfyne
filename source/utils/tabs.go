package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	tabFileSuffix = "Tab.go"
)

// TabNames returns each of the current tab folder names.
// Param buttonName should be an existing button name or there will be an error.
func TabNames(
	buttonName string,
	folderPaths *FolderPaths,
) (tabNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.TabNames: %w", err)
		}
	}()

	var folderNames map[string]string
	if _, folderNames, err = ButtonTabBarFileFolderNames(folderPaths); err != nil {
		return
	}
	var found bool
	var folderName string
	if folderName, found = folderNames[buttonName]; !found {
		err = fmt.Errorf("button named %q, that uses a tab-bar not found", buttonName)
		return
	}
	tabBarFolderPath := filepath.Join(folderPaths.FrontendPanelHome, folderName)
	var maybeTabFiles map[string]string
	var maybeTabFolders map[string]string
	if maybeTabFiles, maybeTabFolders, err = maybeTabFilesFolders(tabBarFolderPath); err != nil {
		return
	}
	tabNames = make([]string, 0, len(maybeTabFiles))
	for tabName := range maybeTabFiles {
		if _, found = maybeTabFolders[tabName]; found {
			tabNames = append(tabNames, tabName)
		}
	}
	return
}

// ValidateNewTabName returns an error if the tab name is not valid or unique.
// Param buttonName should be an existing button name or there will be an error.
func ValidateNewTabName(
	buttonName string,
	tabName string,
	folderPaths *FolderPaths,
	dumperCh chan string,
) (isValid bool, err error) {

	var userMessage string
	if isValid, userMessage = validateName(tabName); !isValid {
		dumperCh <- userMessage
		return
	}

	var tabNames []string
	if tabNames, err = TabNames(buttonName, folderPaths); err != nil {
		return
	}
	for _, name := range tabNames {
		if name == tabName {
			isValid = false
			dumperCh <- fmt.Sprintf("The tab name %q is already being used.", tabName)
			return
		}
	}
	isValid = true
	return
}

// ValidateCurrentTabName returns an error if the tab name is not valid.
// Param buttonName should be an existing button name or there will be an error.
func ValidateCurrentTabName(
	buttonName string,
	tabName string,
	folderPaths *FolderPaths,
) (isValid bool, userMessage string, err error) {

	if isValid, userMessage = validateName(tabName); !isValid {
		return
	}

	var tabNames []string
	if tabNames, err = TabNames(buttonName, folderPaths); err != nil {
		return
	}
	for _, name := range tabNames {
		if name == tabName {
			return
		}
	}
	isValid = false
	userMessage = fmt.Sprintf("Unable to confirm that the button named %q has a tab named %q.", buttonName, tabName)
	return
}

// maybeTabFilesFolders returns tab names mapped to their tab file names and tab names mapped to their panel group folders.
func maybeTabFilesFolders(tabBarFolderPath string) (maybeTabFiles, maybeTabPanelGroupFolders map[string]string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ButtonFileFolderNames: %w", err)
		}
	}()

	// Find all /home/*Button.go files.
	// maybeButtonFiles are the files that might be button files.
	var fNames []string
	var fName string
	var trimmed string
	if fNames, err = FileNames(tabBarFolderPath); err != nil {
		return
	}
	maybeTabFiles = make(map[string]string)
	for _, fName = range fNames {
		if trimmed = strings.TrimSuffix(fName, tabFileSuffix); trimmed != fName {
			if len(fName) == 0 {
				continue
			}
			capped := Cap(trimmed)
			if isValid, _ := validateName(capped); !isValid {
				continue
			}
			// Looks like a tab file.
			maybeTabFiles[capped] = fName
		}
	}

	// Find all folders.
	if fNames, err = FolderNames(tabBarFolderPath); err != nil {
		return
	}
	maybeTabPanelGroupFolders = make(map[string]string)
	for _, fName = range fNames {
		if fName == tabBarFolderSuffix || fName == tabFileSuffix {
			continue
		}
		if trimmed = strings.TrimSuffix(fName, panelGroupFolderSuffix); trimmed != fName {
			// Looks like a kickfyne tab panel-group-folder.
			capped := Cap(trimmed)
			if isValid, _ := validateName(capped); !isValid {
				continue
			}
			maybeTabPanelGroupFolders[capped] = fName
		}
	}
	return
}
