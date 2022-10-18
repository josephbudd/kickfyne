package framework

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd         = "framework"
	verbHelp    = "help"
	verbRestart = "restart"
)

// Handler passes control to the correct handler.
func Handler(pathWD string, args []string, isBuilt bool, importPrefix string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("framework.Handler: %w", err)
		}
	}()

	switch isBuilt {
	case true:
		if len(args) == 0 {
			fmt.Println(Usage())
			return
		}
		var folderPaths *utils.FolderPaths
		if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
			return
		}
		// The framework is built in this folder.
		switch args[0] {
		case verbRestart:
			if err = utils.UnBuildFolderPaths(folderPaths); err != nil {
				return
			}
			fmt.Println("Removed the framework for restart.")
			if err = utils.RebuildFolderPaths(folderPaths); err != nil {
				return
			}
			isBuilt = false
			if err = handleFramework(pathWD, args, isBuilt, importPrefix, "recreated for restart"); err != nil {
				return
			}
		case verbHelp:
			fmt.Println(Usage())
			return
		default:
			fmt.Println(Usage())
			return
		}
	case false:
		// The framework is not built in this folder.
		if len(args) > 0 {
			fmt.Println(Usage())
			return
		}
		if err = handleFramework(pathWD, args, isBuilt, importPrefix, "created"); err != nil {
			return
		}
	}
	fyneAppTOMLRelativeFilePath := utils.FyneAppTOMLRelativeFilePath()
	fmt.Printf("KICKFYNE TODO: %s may need some editing.\n", fyneAppTOMLRelativeFilePath)
	return
}

// handleFramework creates the framework.
func handleFramework(pathWD string, args []string, isBuilt bool, importPrefix string, action string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("framework.handleFramework: %w", err)
		}
	}()

	if isBuilt {
		fmt.Printf("Warning: The app is already built in %q\n", pathWD)
		return
	}
	importBase := filepath.Base(importPrefix)
	currentWD := filepath.Base(pathWD)
	if importBase != currentWD {
		fmt.Println("Warning: You must run kickfyne inside the app's folder.")
		return
	}
	fmt.Printf("Creating the framework in %q.\n", pathWD)
	// Create the framework code.
	var folderPaths *utils.FolderPaths
	if folderPaths, err = utils.BuildFolderPaths(pathWD); err != nil {
		return
	}
	if err = source.CreateFramework(importBase, importPrefix, folderPaths); err != nil {
		return
	}
	fmt.Printf("Success. The framework is %s.\n", action)
	return
}
