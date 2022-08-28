package framework

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd = "framework"
)

func Handler(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {
	if isBuilt {
		dumperCh <- fmt.Sprintf("The app is already built in %q\n", pathWD)
		return
	}
	importBase := filepath.Base(importPrefix)
	currentWD := filepath.Base(pathWD)
	if importBase != currentWD {
		dumperCh <- "Run inside the app's folder."
		return
	}
	dumperCh <- fmt.Sprintf("Building the app in %q\n", pathWD)
	// Build the app code.
	if err = source.Build(importBase, importPrefix, folderPaths); err != nil {
		return
	}
	// Build home with no buttons.
	if err = home.Build(importPrefix, folderPaths); err != nil {
		return
	}
	dumperCh <- "Success. The framework is initialized."
	return
}
