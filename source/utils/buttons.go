package utils

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

const (
	buttonFileSuffix       = "Button.go"
	tabBarFolderSuffix     = "TabBar"
	panelGroupFolderSuffix = "PanelGroup"
)

// HomeButtonFileName returns the file name for a "home/<button name>Button.go" file.
func HomeButtonFileName(buttonName string) (fileName string) {
	fileName = DeCap(buttonName) + buttonFileSuffix
	return
}

// HomeTabBarFileName returns the file name for a "home/<button name>TabBar.go" file
func HomeTabBarFileName(buttonName string) (fileName string) {
	fileName = DeCap(buttonName) + buttonFileSuffix
	return
}

func ButtonTabBarFolderName(buttonName string) (folderName string) {
	folderName = DeCap(buttonName) + tabBarFolderSuffix
	return
}

func ButtonPanelGroupFolderName(buttonName string) (folderName string) {
	folderName = DeCap(buttonName) + panelGroupFolderSuffix
	return
}

func TabPanelGroupFolderName(tabName string) (folderName string) {
	folderName = DeCap(tabName) + panelGroupFolderSuffix
	return
}

func ButtonTabBarFolderImport(importPrefix, buttonName string) (importFolder string) {
	importFolder = path.Join(importPrefix, "/frontend/panel/home/", DeCap(buttonName)+tabBarFolderSuffix)
	return
}

func ButtonPanelGroupFolderImport(importPrefix, buttonName string) (importFolder string) {
	importFolder = path.Join(importPrefix, "/frontend/panel/home/", DeCap(buttonName)+panelGroupFolderSuffix)
	return
}

func ButtonTabBarTabPanelGroupFolderImport(importPrefix, buttonName, tabName string) (importFolder string) {
	importFolder = ButtonTabBarFolderImport(importPrefix, buttonName)
	importFolder = path.Join(importFolder, DeCap(tabName)+panelGroupFolderSuffix)
	return
}

// ButtonHasTabBar returns if the button has a tab bar.
// If false then the button has a panel group.
func ButtonHasTabBar(buttonName string, folderPaths *FolderPaths) (buttonHasTabBar bool, err error) {
	var _, folderPath string
	if _, folderPath, err = ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	folderName := filepath.Base(folderPath)
	if folderName == tabBarFolderSuffix {
		return
	}
	buttonHasTabBar = folderName != strings.TrimSuffix(folderName, tabBarFolderSuffix)
	return
}

// ButtonFileFolderPaths returns a button file and folder paths.
// Param buttonName must be an existing button name or an error will be returned.
func ButtonFileFolderPaths(buttonName string, folderPaths *FolderPaths) (filePath, folderPath string, err error) {

	var maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders map[string]string
	if maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders, err = maybeButtonFilesFolders(folderPaths); err != nil {
		return
	}
	var fileName string
	var folderName string
	var found bool
	// Find the button file.
	if fileName, found = maybeButtonFiles[buttonName]; !found {
		err = fmt.Errorf("button file not found for button named %q", buttonName)
		return
	}
	// Find the button panel group or tab bar folder.
	filePath = filepath.Join(folderPaths.FrontendPanelHome, fileName)
	if folderName, found = maybeButtonPanelGroupFolders[buttonName]; found {
		folderPath = filepath.Join(folderPaths.FrontendPanelHome, folderName)
		return
	}
	if folderName, found = maybeButtonTabBarFolders[buttonName]; found {
		folderPath = filepath.Join(folderPaths.FrontendPanelHome, folderName)
		return
	}
	err = fmt.Errorf("button folder not found for button named %q", buttonName)
	return
}

// ButtonFileFolderNames returns a list of button file and folder names.
func ButtonFileFolderNames(folderPaths *FolderPaths) (verifiedFileNames, verifiedFolderNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ButtonFileFolderNames: %w", err)
		}
	}()

	var maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders map[string]string
	if maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders, err = maybeButtonFilesFolders(folderPaths); err != nil {
		return
	}

	// Assume that maybeButtonFiles keys that are also maybeButtonFolders keys.
	verifiedFileNames = make([]string, 0, len(maybeButtonFiles))
	verifiedFolderNames = make([]string, 0, len(maybeButtonFiles))
	for buttonKey, fileName := range maybeButtonFiles {
		if folderName, found := maybeButtonPanelGroupFolders[buttonKey]; found {
			// This home-button file has a matching panel-group folder.
			verifiedFileNames = append(verifiedFileNames, fileName)
			verifiedFolderNames = append(verifiedFolderNames, folderName)
			continue
		}
		if folderName, found := maybeButtonTabBarFolders[buttonKey]; found {
			// This home-button file has a matching tab-bar folder.
			verifiedFileNames = append(verifiedFileNames, fileName)
			verifiedFolderNames = append(verifiedFolderNames, folderName)
		}
	}
	return
}

// ButtonPanelGroupFileFolderNames returns a list of panel group button file and folder names.
func ButtonPanelGroupFileFolderNames(folderPaths *FolderPaths) (verifiedFileNames, verifiedFolderNames map[string]string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ButtonPanelGroupFileFolderNames: %w", err)
		}
	}()

	var maybeButtonFiles, maybeButtonPanelGroupFolders map[string]string
	if maybeButtonFiles, maybeButtonPanelGroupFolders, _, err = maybeButtonFilesFolders(folderPaths); err != nil {
		return
	}

	// Assume that maybeButtonFiles keys that are also maybeButtonFolders keys.
	verifiedFileNames = make(map[string]string, len(maybeButtonFiles))
	verifiedFolderNames = make(map[string]string, len(maybeButtonFiles))
	for buttonName, fileName := range maybeButtonFiles {
		if folderName, found := maybeButtonPanelGroupFolders[buttonName]; found {
			// This home-button file has a matching panel-group folder.
			verifiedFileNames[buttonName] = fileName
			verifiedFolderNames[buttonName] = folderName
		}
	}
	return
}

// ButtonTabBarFileFolderNames returns a list of tab bar button file and folder names.
func ButtonTabBarFileFolderNames(folderPaths *FolderPaths) (verifiedFileNames, verifiedFolderNames map[string]string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ButtonTabBarFileFolderNames: %w", err)
		}
	}()

	var maybeButtonFiles, maybeButtonTabBarFolders map[string]string
	if maybeButtonFiles, _, maybeButtonTabBarFolders, err = maybeButtonFilesFolders(folderPaths); err != nil {
		return
	}

	// Assume that maybeButtonFiles keys that are also maybeButtonFolders keys.
	verifiedFileNames = make(map[string]string, len(maybeButtonFiles))
	verifiedFolderNames = make(map[string]string, len(maybeButtonFiles))
	for buttonName, fileName := range maybeButtonFiles {
		if folderName, found := maybeButtonTabBarFolders[buttonName]; found {
			// This home-button file has a matching tab-bar folder.
			verifiedFileNames[buttonName] = fileName
			verifiedFolderNames[buttonName] = folderName
		}
	}
	return
}

// ButtonNames returns each of the current button names.
func ButtonNames(folderPaths *FolderPaths) (verifiedButtonNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ButtonNames: %w", err)
		}
	}()

	var maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders map[string]string
	if maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders, err = maybeButtonFilesFolders(folderPaths); err != nil {
		return
	}

	// Assume that maybeButtonFiles keys that are also maybeButtonFolders keys.
	verifiedButtonNames = make([]string, 0, len(maybeButtonFiles))
	for buttonKey := range maybeButtonFiles {
		if _, found := maybeButtonPanelGroupFolders[buttonKey]; found {
			// This home-button file has a matching panel-group folder.
			verifiedButtonNames = append(verifiedButtonNames, buttonKey)
			continue
		}
		if _, found := maybeButtonTabBarFolders[buttonKey]; found {
			// This home-button file has a matching tab-bar folder.
			verifiedButtonNames = append(verifiedButtonNames, buttonKey)
		}
	}
	return
}

// ValidateNewButtonName returns an error if the button name is not valid.
func ValidateNewButtonNames(
	newButtonNames []string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	msgs := make([]string, 0, 10)
	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateNewButtonNames: %w", err)
			return
		}
		if len(msgs) > 0 {
			failureMessage = strings.Join(msgs, "\n")
		}
	}()

	var buttonNames []string
	if buttonNames, err = ButtonNames(folderPaths); err != nil {
		return
	}
	isValid = true
	for _, newButtonName := range newButtonNames {
		// Is the name even a valid button name?
		var isOK bool
		var msg string
		if isOK, msg = validateName(newButtonName); !isOK {
			// Name is not valid. Continue with the next button name.
			isValid = false
			msgs = append(msgs, msg)
			continue
		}
		// Is the name too similar to other button names.
		lcNewButtonName := strings.ToLower(newButtonName)
		for _, name := range buttonNames {
			if name == newButtonName {
				isValid = false
				msg = fmt.Sprintf("The button name %q is already being used.", newButtonName)
				msgs = append(msgs, msg)
				continue
			}
			if strings.ToLower(name) == lcNewButtonName {
				msg = fmt.Sprintf("The button name %q is too much like the button name %q.", name, newButtonName)
				msgs = append(msgs, msg)
				isValid = false
				continue
			}
		}
	}

	return
}

// ValidateNewButtonName returns an error if the button name is not valid.
func ValidateNewButtonName(
	buttonName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateNewButtonName: %w", err)
			return
		}
	}()

	if isValid, failureMessage = validateName(buttonName); !isValid {
		return
	}

	var buttonNames []string
	if buttonNames, err = ButtonNames(folderPaths); err != nil {
		return
	}
	lcButtonName := strings.ToLower(buttonName)
	for _, name := range buttonNames {
		if name == buttonName {
			isValid = false
			failureMessage = fmt.Sprintf("The button name %q is already being used.", buttonName)
			return
		}
		if strings.ToLower(name) == lcButtonName {
			isValid = false
			failureMessage = fmt.Sprintf("The button name %q is too much like the button name %q.", name, buttonName)
			return
		}
		if folderName := strings.ToLower(name); folderName == strings.ToLower(buttonName) {
			isValid = false
			failureMessage = fmt.Sprintf("The button name %q is not being used but the folder name %q is.", buttonName, folderName)
			return
		}
	}
	isValid = true
	return
}

// ValidateCurrentButtonName returns an error if the button name is not valid.
func ValidateCurrentButtonName(
	buttonName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	if isValid, failureMessage = validateName(buttonName); !isValid {
		return
	}

	var buttonNames []string
	if buttonNames, err = ButtonNames(folderPaths); err != nil {
		return
	}
	for _, name := range buttonNames {
		if name == buttonName {
			return
		}
	}
	isValid = false
	failureMessage = fmt.Sprintf("The button name %q is not being used.", buttonName)
	return
}

// maybeButtonFilesFolders returns a list of button file names and button folder names.
func maybeButtonFilesFolders(folderPaths *FolderPaths) (maybeButtonFiles, maybeButtonPanelGroupFolders, maybeButtonTabBarFolders map[string]string, err error) {

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
	if fNames, err = FileNames(folderPaths.FrontendPanelHome); err != nil {
		return
	}
	maybeButtonFiles = make(map[string]string)
	for _, fName = range fNames {
		if trimmed = strings.TrimSuffix(fName, buttonFileSuffix); trimmed != fName {
			if len(trimmed) == 0 {
				return
			}
			// Looks like a home-button file.
			capped := Cap(trimmed)
			if isValid, _ := validateName(capped); isValid {
				maybeButtonFiles[capped] = fName
			}
		}
	}

	// Find all folders.
	if fNames, err = FolderNames(folderPaths.FrontendPanelHome); err != nil {
		return
	}
	maybeButtonPanelGroupFolders = make(map[string]string)
	maybeButtonTabBarFolders = make(map[string]string)
	for _, fName = range fNames {
		if trimmed = strings.TrimSuffix(fName, tabBarFolderSuffix); trimmed != fName {
			if len(trimmed) == 0 {
				continue
			}
			capped := Cap(trimmed)
			if isValid, _ := validateName(capped); !isValid {
				continue
			}
			// Looks like a kickfyne home-button tab-bar-folder.
			maybeButtonTabBarFolders[capped] = fName
			continue
		}
		if trimmed = strings.TrimSuffix(fName, panelGroupFolderSuffix); trimmed != fName {
			if len(trimmed) == 0 {
				continue
			}
			capped := Cap(trimmed)
			if isValid, _ := validateName(capped); !isValid {
				continue
			}
			// Looks like a kickfyne home-button panel-group-folder.
			maybeButtonPanelGroupFolders[capped] = fName
		}
	}
	return
}
