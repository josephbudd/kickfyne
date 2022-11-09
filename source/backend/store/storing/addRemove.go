package storing

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// AddRecord adds a record storing file to shared/store/storing/.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storing.AddRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	storerFilePath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	data := templateData{
		RecordName:     recordName,
		ImportPrefix:   importPrefix,
		StorerFilePath: storerFilePath,
	}
	err = utils.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record storing file from shared/store/storing/.
func RemoveRecord(
	recordName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storing.RemoveRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	err = os.Remove(oPath)
	return
}
