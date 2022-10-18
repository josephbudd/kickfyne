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

	// gui/mainmenu/mainmenu.go
	data := mainMenuTemplateData{
		ImportPrefix: importPrefix,
	}
	oPath := filepath.Join(folderPaths.FrontendGUIMainMenu, mainMenuFileName)
	if err = utils.ProcessTemplate(mainMenuFileName, oPath, mainMenuTemplate, data); err != nil {
		return
	}
	return
}
