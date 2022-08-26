package backpanel

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backpanel.Build: %w", err)
		}
	}()

	var oPath string

	// widget/backpanel/backpanel.go
	oPath = filepath.Join(folderPaths.FrontendWidgetBackPanel, fileName)
	if err = utils.ProcessTemplate(fileName, oPath, template, nil); err != nil {
		return
	}

	return
}
