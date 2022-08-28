package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	folderName = "storer"
	goExt      = ".go"
)

var lExt = len(goExt)

// RecordFileName returns the file name for a record.
func RecordFileName(recordName string) (fileName string) {
	fileName = DeCap(recordName) + ".go"
	return
}

func UserRecordNames(folderPaths *FolderPaths) (recordNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.UserRecordNames: %w", err)
		}
	}()

	recordNames, err = AllRecordNames(folderPaths)
	return
}

func AllRecordNames(folderPaths *FolderPaths) (recordNames []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.AllRecordNames: %w", err)
		}
	}()

	var fileNames []string
	if fileNames, err = FileNames(folderPaths.SharedStoreStorer); err != nil {
		return
	}
	for _, fileName := range fileNames {
		ext := filepath.Ext(fileName)
		if ext != goExt {
			continue
		}
		l := len(fileName) - lExt
		recordName := Cap(fileName[:l])
		recordNames = append(recordNames, recordName)
	}
	return
}

// ValidateNewRecordName returns an error if the record name is not valid.
func ValidateNewRecordName(
	recordName string,
	folderPaths *FolderPaths,
) (isValid bool, userMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateRecordName: %w", err)
		}
	}()

	lc := strings.ToLower(recordName)

	if isValid, userMessage = validateName(recordName); !isValid {
		return
	}

	var recordNames []string
	if recordNames, err = AllRecordNames(folderPaths); err != nil {
		return
	}
	for _, name := range recordNames {
		if strings.ToLower(name) == lc {
			isValid = false
			userMessage = fmt.Sprintf("The record name %q is too much like the record name %q.", recordName, name)
			return
		}
	}
	isValid = true
	return
}

// ValidateCurrentRecordName returns an error if the record name is not valid.
func ValidateCurrentRecordName(
	recordName string,
	folderPaths *FolderPaths,
) (isValid bool, userMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateRecordName: %w", err)
		}
	}()

	if isValid, userMessage = validateName(recordName); !isValid {
		return
	}

	var recordNames []string
	if recordNames, err = UserRecordNames(folderPaths); err != nil {
		return
	}
	for _, name := range recordNames {
		if name == recordName {
			return
		}
	}
	isValid = false
	userMessage = fmt.Sprintf("The record name %q is not being used.", recordName)
	return
}
