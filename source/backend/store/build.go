package store

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the shared/store/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.CreateFramework: %w", err)
		}
	}()

	// stores.go
	if err = createStoresGo(importPrefix, folderPaths); err != nil {
		return
	}

	return
}

// createStoresGo builds the stores.go files with no record names.
func createStoresGo(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.createStoresGo: %w", err)
		}
	}()

	// store/stores.go
	oPath := filepath.Join(folderPaths.SharedStore, storesFileName)
	data := storesTemplateData{
		ImportPrefix: importPrefix,
		RecordNames:  []string{},
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(storesFileName, oPath, storesTemplate, data); err != nil {
		return
	}

	return
}

// rebuildStoresGo builds the stores.go file with all of the record names.
func rebuildStoresGo(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.rebuildStoresGo: %w", err)
		}
	}()

	var recordNames []string
	if recordNames, err = utils.RecordNames(folderPaths); err != nil {
		return
	}

	// store/stores.go
	oPath := filepath.Join(folderPaths.SharedStore, storesFileName)
	data := storesTemplateData{
		ImportPrefix: importPrefix,
		RecordNames:  recordNames,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(storesFileName, oPath, storesTemplate, data); err != nil {
		return
	}

	return
}
