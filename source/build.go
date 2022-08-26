package source

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/source/backend"
	"github.com/josephbudd/kickfyne/source/frontend"
	"github.com/josephbudd/kickfyne/source/shared"
	"github.com/josephbudd/kickfyne/source/utils"
)

func HasAppFolder(currentWP, appName string) (hasAppFolder bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.HasAppFolder: %w", err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(currentWP); err != nil {
		return
	}
	var dirEntry os.DirEntry
	for _, dirEntry = range dirEntrys {
		if dirEntry.IsDir() {
			dName := dirEntry.Name()
			if hasAppFolder = dName == appName; hasAppFolder {
				return
			}
		}
	}
	return
}

// Build builds the framework in an appName folder in this parent folder.
func Build(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.Build: %w", err)
		}
	}()

	// Shared
	if err = shared.BuildShared(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Backend
	if err = backend.Build(importPrefix, folderPaths); err != nil {
		return
	}

	// Frontend
	if err = frontend.Build(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
