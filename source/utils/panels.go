package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

// PanelNames returns each of the current panel names.
func PanelNames(screenPackageFolderPath string) (panelNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelNames: %w", err)
		}
	}()

	var fileNames []string
	fileNames, err = FileNames(screenPackageFolderPath)
	panelNames = make([]string, 0, len(fileNames))
	for _, fileName := range fileNames {
		var trimmed string
		if fileName == panelFileSuffix {
			// This is not a panel file name.
			continue
		}
		trimmed = strings.TrimSuffix(fileName, panelFileSuffix)
		if trimmed != fileName {
			// This file name has the panel file name suffix.
			// This might be a panel file name.
			panelName := strings.TrimSuffix(fileName, goFileExt)
			panelNames = append(panelNames, panelName)
		}
	}
	return
}

func IsFrameWorkPanelName(panelName string) (is bool) {
	switch panelName {
	case AccordionPanelName:
		is = true
	case AppTabsPanelName:
		is = true
	case DocTabsPanelName:
		is = true
	}
	return
}

// ValidateNewScreenPanelName validates a new panel name for a screenName.
func ValidateNewScreenPanelName(
	screenPackageName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateNewScreenPanelName: %w", err)
		}
	}()

	if isValid = !IsFrameWorkPanelName(panelName); !isValid {
		failureMessage = fmt.Sprintf("The panel name %q is reserved for the framework.", panelName)
		return
	}

	if isValid, failureMessage = validateScreenPanelName(panelName); !isValid {
		return
	}

	screenPackageFolderPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName)
	var currentPanelNames []string
	if currentPanelNames, err = PanelNames(screenPackageFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if panelName == currentPanelName {
			// This is not a new panel name.
			isValid = false
			failureMessage = fmt.Sprintf("The screen package %q already has a panel named %q.", screenPackageName, panelName)
			return
		}
	}
	// This is a new panel name.
	isValid = true
	return
}

// ValidateCurrentScreenPanelName validates a new panel name for a screenName.
func ValidateCurrentScreenPanelName(
	screenPackageName string,
	panelName string,
	folderPaths *FolderPaths,
) (isValid bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("ValidateCurrentScreenPanelName: %w", err)
		}
	}()

	if isValid = !IsFrameWorkPanelName(panelName); !isValid {
		failureMessage = fmt.Sprintf("The panel name %q is reserved for the framework.", panelName)
		return
	}

	if isValid, failureMessage = validateScreenPanelName(panelName); !isValid {
		return
	}

	screenPackageFolderPath := filepath.Join(folderPaths.FrontendGUIScreens, screenPackageName)
	var currentPanelNames []string
	if currentPanelNames, err = PanelNames(screenPackageFolderPath); err != nil {
		return
	}
	for _, currentPanelName := range currentPanelNames {
		if panelName == currentPanelName {
			// this is a current panel name.
			isValid = true
			return
		}
	}
	// This is not a current panel name.
	log.Printf("currentPanelNames %#v", currentPanelNames)
	failureMessage = fmt.Sprintf("The screen package %q does not have a panel named %q.", screenPackageName, panelName)
	return
}
