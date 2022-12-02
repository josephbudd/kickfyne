package mainmenu

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// Build creates the framework's frontend/gui/mainmenu/ files.
func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.Build: %w", err)
		}
	}()

	var screenPackageNames []string
	if screenPackageNames, err = utils.ScreenPackageNames(folderPaths); err != nil {
		return
	}
	// gui/mainmenu/mainmenu.go
	data := mainMenuTemplateData{
		ImportPrefix:       importPrefix,
		ScreenPackageNames: screenPackageNames,
		HomePackageName:    utils.HomeScreenPackageName,
		Funcs:              utils.GetFuncs(),
	}
	oPath := filepath.Join(folderPaths.FrontendGUIMainMenu, mainMenuFileName)
	if err = utils.ProcessTemplate(mainMenuFileName, oPath, mainMenuTemplate, data); err != nil {
		return
	}
	return
}
