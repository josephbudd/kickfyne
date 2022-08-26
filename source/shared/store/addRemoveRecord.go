package store

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/shared/store/record"
	"github.com/josephbudd/kickfyne/source/shared/store/storer"
	"github.com/josephbudd/kickfyne/source/shared/store/storing"
	"github.com/josephbudd/kickfyne/source/utils"
)

// AddRecord add the files for the new record and then rebuilds stores.go.
func AddRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.AddRecord: %w", err)
		}
	}()

	if err = record.AddRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = storer.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	if err = storing.AddRecord(recordName, importPrefix, folderPaths); err != nil {
		return
	}
	err = rebuildStoresGo(importPrefix, folderPaths)
	return
}

// RemoveRecord add the files for the new record and then rebuilds stores.go.
func RemoveRecord(
	recordName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.RemoveRecord: %w", err)
		}
	}()

	if err = record.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = storer.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	if err = storing.RemoveRecord(recordName, folderPaths); err != nil {
		return
	}
	err = rebuildStoresGo(importPrefix, folderPaths)
	return
}
