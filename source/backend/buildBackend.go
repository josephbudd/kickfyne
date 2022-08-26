package backend

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/backend/misc/shuffle"
	"github.com/josephbudd/kickfyne/source/backend/txrx"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "backend"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backend.Build: %w", err)
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

	// backend/misc/shuffle/
	if err = shuffle.Build(folderPaths); err != nil {
		return
	}

	// backend/txrx/
	if err = txrx.Build(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
