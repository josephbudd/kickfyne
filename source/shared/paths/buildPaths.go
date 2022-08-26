package paths

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	appName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("paths.Build: %w", err)
		}
	}()

	// paths/paths.go
	oPath := filepath.Join(folderPaths.SharedPaths, fileName)
	data := templateData{
		AppName: appName,
		Funcs:   utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(fileName, oPath, template, data)
	return
}
