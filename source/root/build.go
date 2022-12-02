package root

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the shared/ files.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shared.CreateFramework: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// ./FyneApp.toml
	importPrefixParts := strings.Split(importPrefix, "/")
	lIPP := len(importPrefixParts)
	appIDParts := make([]string, 0, lIPP*2)
	for _, importPrefixPart := range importPrefixParts {
		parts := strings.Split(importPrefixPart, ".")
		for i := len(parts) - 1; i >= 0; i-- {
			appIDParts = append(appIDParts, parts[i])
		}
	}

	data = fyneAppTOMLData{
		WebSiteURL:      "https://" + importPrefix + "/",
		AppName:         appName,
		AppID:           strings.Join(appIDParts, "."),
		HomePackageName: utils.HomeScreenPackageName,
	}
	if err = utils.ProcessTemplate(utils.FyneAppTOMLFileName, utils.FyneAppTOMLFilePath(folderPaths), dyneAppTOMLTemplate, data); err != nil {
		return
	}

	var screenNames []string
	if screenNames, err = utils.ScreenPackageNames(folderPaths); err != nil {
		return
	}

	// ./main.go
	oPath = filepath.Join(folderPaths.App, MainFileName)
	data = mainTemplateData{
		ImportPrefix:       importPrefix,
		AppName:            appName,
		ScreenPackageNames: screenNames,
		HomePackageName:    utils.HomeScreenPackageName,
		Funcs:              utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(MainFileName, oPath, mainTemplate, data); err != nil {
		return
	}
	if err = RebuildMainScreensGo(
		appName,
		screenNames,
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	// vscode workspaces
	oPath = filepath.Join(folderPaths.App, backendWorkSpaceFileName)
	if err = utils.ProcessTemplate(backendWorkSpaceFileName, oPath, backendWorkSpaceTemplate, nil); err != nil {
		return
	}
	oPath = filepath.Join(folderPaths.App, frontendWorkSpaceFileName)
	if err = utils.ProcessTemplate(frontendWorkSpaceFileName, oPath, frontendWorkSpaceTemplate, nil); err != nil {
		return
	}

	return
}

func RebuildMainScreensGo(
	appName string,
	sortedScreenNames []string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {
	oPath := filepath.Join(folderPaths.App, ScreensFileName)
	data := screensTemplateData{
		ImportPrefix:       importPrefix,
		AppName:            appName,
		ScreenPackageNames: sortedScreenNames,
		HomePackageName:    utils.HomeScreenPackageName,
		Funcs:              utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(ScreensFileName, oPath, screensTemplate, data)
	return
}
