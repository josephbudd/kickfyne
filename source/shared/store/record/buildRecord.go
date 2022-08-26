package record

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
			err = fmt.Errorf("record.Build: %w", err)
		}
	}()

	oPath := filepath.Join(folderPaths.SharedStoreRecord, stateFileName)
	err = utils.ProcessTemplate(stateFileName, oPath, stateTemplate, nil)

	return
}
