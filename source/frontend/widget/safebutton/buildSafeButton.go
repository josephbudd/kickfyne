package safebutton

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
			err = fmt.Errorf("safebutton.Build: %w", err)
		}
	}()

	// widget/safebutton/safebutton.go
	oPath := filepath.Join(folderPaths.FrontendWidgetSafeButton, fileName)
	if err = utils.ProcessTemplate(fileName, oPath, template, nil); err != nil {
		return
	}

	return
}
