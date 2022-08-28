package shared

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/shared/message"
	"github.com/josephbudd/kickfyne/source/shared/paths"
	"github.com/josephbudd/kickfyne/source/shared/store"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "shared"
)

func BuildShared(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shared.BuildShared: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// shared/main.go
	oPath = filepath.Join(folderPaths.Shared, MainFileName)
	data = mainTemplateData{
		ImportPrefix: importPrefix,
		AppName:      appName,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(MainFileName, oPath, mainTemplate, data); err != nil {
		return
	}

	// shared/message/
	if err = message.Build(folderPaths); err != nil {
		return
	}

	// shared/paths/
	if err = paths.Build(appName, folderPaths); err != nil {
		return
	}

	// shared/store/
	if err = store.Build(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
