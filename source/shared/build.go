package shared

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/shared/message"
	"github.com/josephbudd/kickfyne/source/shared/metadata"
	"github.com/josephbudd/kickfyne/source/shared/paths"
	"github.com/josephbudd/kickfyne/source/shared/store"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "shared"
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

	// shared/shared.go
	path := filepath.Join(folderPaths.Shared, sharedFileName)
	data := sharedTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(sharedFileName, path, sharedTemplate, data); err != nil {
		return
	}

	// shared/message/
	if err = message.CreateFramework(folderPaths); err != nil {
		return
	}

	// shared/metadata/
	if err = metadata.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// shared/paths/
	if err = paths.CreateFramework(appName, folderPaths); err != nil {
		return
	}

	// shared/store/
	if err = store.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
