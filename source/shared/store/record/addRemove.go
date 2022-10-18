package record

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// AddRecord adds a record file to shared/store/record/.
func AddRecord(
	recordName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.AddRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreRecord, fName)
	data := templateData{
		RecordName: recordName,
		Funcs:      utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record file from shared/store/record/.
func RemoveRecord(
	recordName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.RemoveRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.SharedStoreRecord, fName)
	err = os.Remove(oPath)
	return
}
