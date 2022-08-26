package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	PanelFileSuffix = "Panel.go"
)

// PanelNames returns each of the current panel folder names.
func PanelNames(panelGroupFolderPath string) (panelNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("PanelNames: %w", err)
		}
	}()

	var fileNames []string
	fileNames, err = FileNames(panelGroupFolderPath)
	panelNames = make([]string, 0, len(fileNames))
	for _, fileName := range fileNames {
		var trimmed string
		if fileName == PanelFileSuffix {
			// Not a panel file.
			continue
		}
		// Does this panel file name use a valid panel name.
		trimmed = strings.TrimSuffix(fileName, PanelFileSuffix)
		if trimmed != fileName {
			panelNames = append(panelNames, Cap(trimmed))
		}
	}
	return
}

// ValidateNewButtonPanelName validates a new panel name for a button's panel group.
func ValidateNewButtonPanelName(
	buttonName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateNewButtonPanelName: %w", err)
		}
	}()

	var buttonFolderPath string
	if _, buttonFolderPath, err = ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	buttonHasPanelGroup := filepath.Base(buttonFolderPath) == ButtonPanelGroupFolderName(buttonName)
	var currentPanelNames []string
	// This panel is to be added to a button panel group.
	if !buttonHasPanelGroup {
		failureMessage = fmt.Sprintf("The button named %q has tabs not panels.", buttonName)
		return
	}
	// The panel name must be new to this button's panel group.
	if currentPanelNames, err = PanelNames(buttonFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if currentPanelName == panelName {
			failureMessage = fmt.Sprintf("The button named %q already has a panel named %q", buttonName, panelName)
			return
		}
	}
	// Not a current button panel name.
	isValid = true
	return
}

// ValidateCurrentButtonPanelName validates a new panel name for a button's panel group.
func ValidateCurrentButtonPanelName(
	buttonName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateCurrentButtonPanelName: %w", err)
		}
	}()

	var buttonFolderPath string
	if _, buttonFolderPath, err = ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	buttonHasPanelGroup := filepath.Base(buttonFolderPath) == ButtonPanelGroupFolderName(buttonName)
	var currentPanelNames []string
	// This panel is to be added to a button panel group.
	if !buttonHasPanelGroup {
		failureMessage = fmt.Sprintf("The button named %q has tabs not panels.", buttonName)
		return
	}
	// The panel name must be new to this button's panel group.
	if currentPanelNames, err = PanelNames(buttonFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if currentPanelName == panelName {
			isValid = true
			return
		}
	}
	// Not a current button panel name.
	failureMessage = fmt.Sprintf("The button named %q does not have a panel named %q", buttonName, panelName)
	return
}

// ValidateNewTabPanelName validates a new panel name for a tab's panel group.
func ValidateNewTabPanelName(
	buttonName string,
	tabName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateNewTabPanelName: %w", err)
		}
	}()

	var buttonFolderPath string
	if _, buttonFolderPath, err = ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	buttonHasPanelGroup := filepath.Base(buttonFolderPath) == ButtonPanelGroupFolderName(buttonName)
	// This panel is to be added to a tab bar panel group.
	if buttonHasPanelGroup {
		failureMessage = fmt.Sprintf("The button named %q has panels not tabs.", buttonName)
		return
	}
	// The panel name must be new to this panel group.
	tabPanelGroupFolderName := TabPanelGroupFolderName(tabName)
	tabPanelGroupFolderPath := filepath.Join(buttonFolderPath, tabPanelGroupFolderName)
	// The panel name must be new to this tab's panel group.
	var currentPanelNames []string
	if currentPanelNames, err = PanelNames(tabPanelGroupFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if currentPanelName == panelName {
			failureMessage = fmt.Sprintf("The tab named %q already has a panel named %q", tabName, panelName)
			return
		}
	}
	// Not a current tab panel name.
	isValid = true
	return
}

// ValidateCurrentTabPanelName validates a new panel name for a tab's panel group.
func ValidateCurrentTabPanelName(
	buttonName string,
	tabName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateCurrentTabPanelName: %w", err)
		}
	}()

	var buttonFolderPath string
	if _, buttonFolderPath, err = ButtonFileFolderPaths(buttonName, folderPaths); err != nil {
		return
	}
	buttonHasPanelGroup := filepath.Base(buttonFolderPath) == ButtonPanelGroupFolderName(buttonName)
	// This panel is to be added to a tab bar panel group.
	if buttonHasPanelGroup {
		failureMessage = fmt.Sprintf("The button named %q has panels not tabs.", buttonName)
		return
	}
	// The panel name must be new to this panel group.
	tabPanelGroupFolderName := TabPanelGroupFolderName(tabName)
	tabPanelGroupFolderPath := filepath.Join(buttonFolderPath, tabPanelGroupFolderName)
	// The panel name must be new to this tab's panel group.
	var currentPanelNames []string
	if currentPanelNames, err = PanelNames(tabPanelGroupFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if currentPanelName == panelName {
			isValid = true
			return
		}
	}
	// Not a current tab panel name.
	failureMessage = fmt.Sprintf("The tab named %q does not have a panel named %q.", tabName, panelName)
	return
}
