package source

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/source/backend"
	"github.com/josephbudd/kickfyne/source/frontend"
	"github.com/josephbudd/kickfyne/source/root"
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

// CreateFramework builds the framework in an appName folder in this parent folder.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.CreateFramework: %w", err)
		}
	}()

	// App folder.
	if err = root.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Shared
	if err = shared.CreateFramework(appName, importPrefix, folderPaths); err != nil {
		return
	}

	// Backend
	if err = backend.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// Frontend
	if err = frontend.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
