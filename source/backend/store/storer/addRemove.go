package storer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// AddRecord adds a record storer file to shared/store/storer/.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storer.AddRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.BackendStoreStorer, fName)
	storingFilePath := filepath.Join(folderPaths.BackendStoreStoring, fName)
	data := templateData{
		RecordName:      recordName,
		ImportPrefix:    importPrefix,
		StoringFilePath: storingFilePath,
	}
	err = utils.ProcessTemplate(fName, oPath, template, data)
	return
}

// RemoveRecord removes a record storer file from shared/store/storer/.
func RemoveRecord(
	recordName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("storer.RemoveRecord: %w", err)
		}
	}()

	fName := utils.RecordFileName(recordName)
	oPath := filepath.Join(folderPaths.BackendStoreStorer, fName)
	err = os.Remove(oPath)
	return
}
