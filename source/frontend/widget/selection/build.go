package selection

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the select widgets for the framework.
func CreateFramework(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("selection.CreateFramework: %w", err)
		}
	}()

	// widget/selection/checkGroup.go
	oPath := filepath.Join(folderPaths.FrontendWidgetSelection, checkGroupFileName)
	if err = utils.ProcessTemplate(checkGroupFileName, oPath, checkGroupTemplate, nil); err != nil {
		return
	}

	// widget/selection/radioGroup.go
	oPath = filepath.Join(folderPaths.FrontendWidgetSelection, radioGroupFileName)
	if err = utils.ProcessTemplate(radioGroupFileName, oPath, radioGroupTemplate, nil); err != nil {
		return
	}

	return
}
