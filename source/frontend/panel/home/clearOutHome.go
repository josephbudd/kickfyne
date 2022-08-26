package home

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func ClearOut(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home.Demolish: %w", err)
		}
	}()

	// Get the button file and folder names.
	var fileNames []string
	var folderNames []string
	if fileNames, folderNames, err = utils.ButtonFileFolderNames(folderPaths); err != nil {
		return
	}

	var fName string
	var path string
	for _, fName = range fileNames {
		// Remove the home-button file.
		path = filepath.Join(folderPaths.FrontendPanelHome, fName)
		if err = os.Remove(path); err != nil {
			return
		}
	}
	for _, fName = range folderNames {
		// Remove the folder.
		path = filepath.Join(folderPaths.FrontendPanelHome, fName)
		if err = os.RemoveAll(path); err != nil {
			return
		}
	}
	return
}
