package gui

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/gui/mainmenu"
	"github.com/josephbudd/kickfyne/source/frontend/gui/screens"
	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/gui/ file.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("gui.CreateFramework: %w", err)
		}
	}()

	// gui/gui.go
	oPath := filepath.Join(folderPaths.FrontendGUI, guiFileName)
	if err = utils.ProcessTemplate(guiFileName, oPath, guiTemplate, nil); err != nil {
		return
	}

	// gui/canvasObjectProvider.go
	oPath = filepath.Join(folderPaths.FrontendGUI, canvasObjectProviderFileName)
	if err = utils.ProcessTemplate(canvasObjectProviderFileName, oPath, canvasObjectProviderTemplate, nil); err != nil {
		return
	}

	// gui/windowWatcher.go
	oPath = filepath.Join(folderPaths.FrontendGUI, windowWatcherFileName)
	if err = utils.ProcessTemplate(windowWatcherFileName, oPath, windowWatcherTemplate, nil); err != nil {
		return
	}

	// gui/accordionItemWatcher.go
	oPath = filepath.Join(folderPaths.FrontendGUI, accordionItemWatcherFileName)
	if err = utils.ProcessTemplate(accordionItemWatcherFileName, oPath, accordionItemWatcherTemplate, nil); err != nil {
		return
	}

	// gui/tabItemWatcher.go
	oPath = filepath.Join(folderPaths.FrontendGUI, tabItemWatcherFileName)
	if err = utils.ProcessTemplate(tabItemWatcherFileName, oPath, tabItemWatcherTemplate, nil); err != nil {
		return
	}

	// Get the name of the landing screen from the meta data.
	var metaData utils.FyneAppMetaData
	if metaData, err = utils.ReadMetaData(folderPaths); err != nil {
		return
	}
	// gui/screens/<metaData.FrontEnd.Landing>/ package with no panels.
	if err = screens.BuildPackageWithoutPanels(
		metaData.FrontEnd.Landing,
		fmt.Sprintf("Package %s is the application's default landing screen.", metaData.FrontEnd.Landing),
		importPrefix,
		folderPaths,
	); err != nil {
		return
	}

	// gui/mainmenu/ package
	err = mainmenu.Build(importPrefix, folderPaths)
	return
}
