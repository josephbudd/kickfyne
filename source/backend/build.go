package backend

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/backend/folder"
	"github.com/josephbudd/kickfyne/source/backend/store"
	"github.com/josephbudd/kickfyne/source/backend/txrx"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "backend"
)

// CreateFramework creates the backend/ files.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backend.CreateFramework: %w", err)
		}
	}()

	// backend/folder/
	if err = folder.CreateFramework(appName, folderPaths); err != nil {
		return
	}

	// backend/store/
	if err = store.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// backend/txrx/
	if err = txrx.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
