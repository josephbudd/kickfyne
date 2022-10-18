package utils

import (
	"path"
	"path/filepath"
)

const (
	panelFileSuffix     = "Panel.go"
	goFileExt           = ".go"
	FyneAppTOMLFileName = "FyneApp.toml"
	ScreenFileName      = "screen.go"
	PanelNameSuffix     = "Panel"
	DocFileName         = "doc.go"

	AccordionPanelFileName = "accordionPanel.go"
	AccordionPanelName     = "accordionPanel"

	AppTabsPanelFileName = "apptabsPanel.go"
	AppTabsPanelName     = "apptabsPanel"

	DocTabsPanelFileName = "doctabsPanel.go"
	DocTabsPanelName     = "doctabsPanel"
)

// FyneAppTOMLFilePath
func FyneAppTOMLFilePath(folderPaths *FolderPaths) (metaDataTOMLFilePath string) {
	metaDataTOMLFilePath = filepath.Join(folderPaths.App, FyneAppTOMLFileName)
	return
}

// FyneAppTOMLRelativeFilePath returns the relative path for the FyneApp.toml file.
func FyneAppTOMLRelativeFilePath() (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameShared, FyneAppTOMLFileName)
	return
}

// PanelFileName returns the file name for a panel file.
func PanelFileName(panelName string) (fileName string) {
	fileName = panelName + goFileExt
	return
}

// PanelFileRelativeFilePath returns the relative path for a panel file.
func PanelFileRelativeFilePath(screenPackageName, panelName string) (relativeFilePath string) {
	fileName := panelName + goFileExt
	relativeFilePath = path.Join(folderNameFrontend, folderNameGUI, FolderNameScreens, screenPackageName, fileName)
	return
}

// ScreenFileRelativeFilePath returns the relative path for a screen's screen.go file.
func ScreenFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, folderNameGUI, FolderNameScreens, screenPackageName, ScreenFileName)
	return
}

// DocFileRelativeFilePath returns the relative path for a screen's screen.go file.
func DocFileRelativeFilePath(screenPackageName string) (relativeFilePath string) {
	relativeFilePath = path.Join(folderNameFrontend, folderNameGUI, FolderNameScreens, screenPackageName, DocFileName)
	return
}

// MessageFileName returns the file name for a messsage.
func MessageFileName(messageName string) (fileName string) {
	fileName = DeCap(messageName) + goFileExt
	return
}

// RecordFileName returns the file name for a record.
func RecordFileName(recordName string) (fileName string) {
	fileName = DeCap(recordName) + goFileExt
	return
}
