package source

import (
	"fmt"

	bestore "github.com/josephbudd/kickfyne/source/backend/store"
	shstore "github.com/josephbudd/kickfyne/source/shared/store"
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

	if err = bestore.RemoveRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	if err = shstore.RemoveRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}

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

	if err = bestore.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	if err = shstore.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}

	return
}
