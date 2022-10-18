package source

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/shared/store"
	"github.com/josephbudd/kickfyne/source/utils"
)

// RemoveRecord removes a record.
func RemoveRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.RemoveRecord: %w", err)
		}
	}()

	err = store.RemoveRecord(recordName, importPrefix, folderPaths)

	return
}

// AddRecord adds a record.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.AddRecord: %w", err)
		}
	}()

	err = store.AddRecord(recordName, importPrefix, folderPaths)

	return
}
