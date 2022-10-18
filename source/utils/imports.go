package utils

import "path"

const (
	pathSeperator = "/"
)

var frontendGUILandingImportSubPath = path.Join(folderNameFrontend, folderNameGUI, FolderNameScreens)

// ScreenFolderImport returns a screens package import path.
func ScreenFolderImport(importPrefix, screenPackageName string) (importFolder string) {
	importFolder = path.Join(importPrefix, frontendGUILandingImportSubPath, screenPackageName)
	return
}
