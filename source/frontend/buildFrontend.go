package frontend

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/panel"
	"github.com/josephbudd/kickfyne/source/frontend/panel/builder"
	"github.com/josephbudd/kickfyne/source/frontend/txrx"
	"github.com/josephbudd/kickfyne/source/frontend/widget"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Build: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// frontend/frontend.go
	oPath = filepath.Join(folderPaths.Frontend, frontendFileName)
	data = frontendTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(frontendFileName, oPath, frontendTemplate, data); err != nil {
		return
	}

	// frontend/panel/
	if err = panel.Build(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/panel/builder
	if err = builder.Build(folderPaths); err != nil {
		return
	}

	// frontend/widget/
	if err = widget.Build(folderPaths); err != nil {
		return
	}

	// frontend/txrx/
	err = txrx.Build(importPrefix, folderPaths)
	return
}
