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
		WebSiteURL: "https://" + importPrefix + "/",
		AppName:    appName,
		AppID:      strings.Join(appIDParts, "."),
	}
	if err = utils.ProcessTemplate(utils.FyneAppTOMLFileName, utils.FyneAppTOMLFilePath(folderPaths), dyneAppTOMLTemplate, data); err != nil {
		return
	}

	// ./main.go
	oPath = filepath.Join(folderPaths.App, MainFileName)
	data = mainTemplateData{
		ImportPrefix: importPrefix,
		AppName:      appName,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(MainFileName, oPath, mainTemplate, data); err != nil {
		return
	}

	return
}
