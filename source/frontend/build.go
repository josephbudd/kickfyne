package frontend

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/frontend/gui"
	"github.com/josephbudd/kickfyne/source/frontend/landingscreen"
	"github.com/josephbudd/kickfyne/source/frontend/txrx"
	"github.com/josephbudd/kickfyne/source/frontend/widget"
	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's frontend/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.CreateFramework: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// frontend/landing.go
	if err = landingscreen.BuildLanding(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/frontend.go
	oPath = filepath.Join(folderPaths.Frontend, frontendFileName)
	data = frontendTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(frontendFileName, oPath, frontendTemplate, data); err != nil {
		return
	}

	// frontend/gui/
	if err = gui.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	// frontend/widget/
	if err = widget.CreateFramework(folderPaths); err != nil {
		return
	}

	// frontend/txrx/
	err = txrx.CreateFramework(importPrefix, folderPaths)
	return
}
