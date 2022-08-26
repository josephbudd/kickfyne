package store

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/source/shared/store/record"
	"github.com/josephbudd/kickfyne/source/shared/store/storer"
	"github.com/josephbudd/kickfyne/source/shared/store/storing"
	"github.com/josephbudd/kickfyne/source/utils"
)

var defaultRecordNames = []string{utils.StateRecordName}

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.Build: %w", err)
		}
	}()

	// stores.go
	if err = buildStoresGo(importPrefix, folderPaths); err != nil {
		return
	}

	// shared/store/record/
	if err = record.Build(folderPaths); err != nil {
		return
	}

	// shared/store/storer/
	if err = storer.Build(importPrefix, folderPaths); err != nil {
		return
	}

	// shared/store/storing/
	if err = storing.Build(importPrefix, folderPaths); err != nil {
		return
	}

	return
}

// buildStoresGo builds the stores.go files with no record names.
func buildStoresGo(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.buildStoresGo: %w", err)
		}
	}()

	// store/stores.go
	oPath := filepath.Join(folderPaths.SharedStore, storesFileName)
	paddedRecordNames := padRecordNames(defaultRecordNames)
	paddedColonRecordNames := padColonRecordNames(defaultRecordNames)
	data := storesTemplateData{
		ImportPrefix:           importPrefix,
		RecordNamesPadded:      paddedRecordNames,      // recordNamesPadded,
		RecordNamesColonPadded: paddedColonRecordNames, // recordNamesColonPadded,
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
	if recordNames, err = utils.AllRecordNames(folderPaths); err != nil {
		return
	}

	// store/stores.go
	oPath := filepath.Join(folderPaths.SharedStore, storesFileName)
	paddedRecordNames := padRecordNames(recordNames)
	paddedColonRecordNames := padColonRecordNames(recordNames)
	data := storesTemplateData{
		ImportPrefix:           importPrefix,
		RecordNamesPadded:      paddedRecordNames,      // recordNamesPadded,
		RecordNamesColonPadded: paddedColonRecordNames, // recordNamesColonPadded,
	}
	if err = utils.ProcessTemplate(storesFileName, oPath, storesTemplate, data); err != nil {
		return
	}

	return
}

func padRecordNames(recordNames []string) (padded map[string]string) {
	var longest int
	lenRecordNames := len(recordNames)
	nameLengths := make([]int, lenRecordNames)
	for i, name := range recordNames {
		nameLength := len(name)
		nameLengths[i] = nameLength
		if nameLength > longest {
			longest = nameLength
		}
	}
	padded = make(map[string]string, lenRecordNames)
	var builder strings.Builder
	for i, name := range recordNames {
		builder.WriteString(name)
		builder.WriteString(" ")
		paddingLength := longest - nameLengths[i]
		for j := paddingLength; j > 0; j-- {
			builder.WriteString(" ")
		}
		padded[name] = builder.String()
		builder.Reset()
	}
	return
}

func padColonRecordNames(recordNames []string) (padded map[string]string) {
	var longest int
	lenRecordNames := len(recordNames)
	nameLengths := make([]int, lenRecordNames)
	for i, name := range recordNames {
		nameLength := len(name)
		nameLengths[i] = nameLength
		if nameLength > longest {
			longest = nameLength
		}
	}
	padded = make(map[string]string, lenRecordNames)
	var builder strings.Builder
	for i, name := range recordNames {
		builder.WriteString(name)
		builder.WriteString(":")
		paddingLength := longest - nameLengths[i]
		for j := paddingLength; j > 0; j-- {
			builder.WriteString(" ")
		}
		padded[name] = builder.String()
		builder.Reset()
	}
	return
}
