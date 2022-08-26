package storer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

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
	oPath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	storingFilePath := filepath.Join(folderPaths.SharedStoreStoring, fName)
	data := templateData{
		RecordName:      recordName,
		ImportPrefix:    importPrefix,
		StoringFilePath: storingFilePath,
	}
	err = utils.ProcessTemplate(fName, oPath, template, data)
	return
}

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
	oPath := filepath.Join(folderPaths.SharedStoreStorer, fName)
	err = os.Remove(oPath)
	return
}
