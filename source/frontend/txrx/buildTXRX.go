package txrx

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.Build: %w", err)
		}
	}()

	// frontend/txrx/listen.go
	oPath := filepath.Join(folderPaths.FrontendTXRX, fileName)
	data := templateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(fileName, oPath, template, data); err != nil {
		return
	}

	return
}
