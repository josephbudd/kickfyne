package selection

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
			err = fmt.Errorf("selection.Build: %w", err)
		}
	}()

	// widget/selection/checkGroup.go
	oPath := filepath.Join(folderPaths.FrontendWidgetSelection, checkGrouopFileName)
	if err = utils.ProcessTemplate(checkGrouopFileName, oPath, checkGroupTemplate, nil); err != nil {
		return
	}

	return
}
