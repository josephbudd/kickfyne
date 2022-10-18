package backend

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/backend/txrx"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "backend"
)

// CreateFramework creates the backend/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backend.CreateFramework: %w", err)
		}
	}()

	// backend/backend.go
	oPath := filepath.Join(folderPaths.Backend, fileName)
	data := templateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(fileName, oPath, template, data); err != nil {
		return
	}

	// backend/txrx/
	if err = txrx.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
