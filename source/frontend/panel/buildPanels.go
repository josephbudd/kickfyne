package panel

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/panel/builder"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panel.Build: %w", err)
		}
	}()

	var oPath string

	// panel/doc.go
	oPath = filepath.Join(folderPaths.FrontendPanel, docFileName)
	if err = utils.ProcessTemplate(docFileName, oPath, docTemplate, nil); err != nil {
		return
	}

	// panel/groupid.go
	oPath = filepath.Join(folderPaths.FrontendPanel, groupIDFileName)
	if err = utils.ProcessTemplate(groupIDFileName, oPath, groupIDTemplate, nil); err != nil {
		return
	}

	// panel/builder
	if err = builder.Build(folderPaths); err != nil {
		return
	}

	return
}
