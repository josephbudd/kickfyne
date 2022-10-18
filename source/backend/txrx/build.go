package txrx

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the framework's backend/txrx/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("txrx.CreateFramework: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// txrx/doc.go
	oPath = filepath.Join(folderPaths.BackendTXRX, docFileName)
	if err = utils.ProcessTemplate(docFileName, oPath, docTemplate, nil); err != nil {
		return
	}

	// txrx/txrx.go
	oPath = filepath.Join(folderPaths.BackendTXRX, tXRXFileName)
	data = tXRXTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(tXRXFileName, oPath, tXRXTemplate, data); err != nil {
		return
	}

	// txrx/Init.go
	fname := utils.MessageFileName(utils.InitMessageName)
	oPath = filepath.Join(folderPaths.BackendTXRX, fname)
	data = initRXTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(fname, oPath, initRXTemplate, data); err != nil {
		return
	}

	return
}
