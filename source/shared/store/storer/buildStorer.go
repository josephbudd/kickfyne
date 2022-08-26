package storer

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
			err = fmt.Errorf("record.Build: %w", err)
		}
	}()

	oPath := filepath.Join(folderPaths.SharedStoreStorer, stateFileName)
	data := stateTemplateData{
		ImportPrefix: importPrefix,
	}
	err = utils.ProcessTemplate(stateFileName, oPath, stateTemplate, data)

	return
}
