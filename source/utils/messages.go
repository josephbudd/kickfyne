package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	InitMessageName          = "Init"
	chansMessageName         = "Chans"
	initMessageNameInvalidF  = `The message name %q is too similar to message name "Init" which belongs to the framework, is used for the application initialization and can not be removed.`
	chansMessageNameInvalidF = `The message name %q is too similar to the file name "chans.go" which belongs to the framework, defines message channels for the application and can not be removed.`
	initMessageNameInvalid   = `The message name "Init" belongs to the framework, is used for the application initialization and can not be removed.`
	chansMessageNameInvalid  = `The message name "Chans" is too much like the file name "chans.go" which belongs to the framework, defines message channels for the application and can not be removed.`
)

// UserMessageNames returns each of the user added message names.
func UserMessageNames(folderPaths *FolderPaths) (names []string, err error) {
	var allNames []string
	if allNames, err = AllMessageNames(folderPaths); err != nil {
		return
	}
	names = make([]string, 0, len(allNames)-1)
	for _, name := range allNames {
		switch {
		case name == InitMessageName:
			continue
		case name == chansMessageName:
			continue
		default:
			names = append(names, name)
		}
	}
	return
}

// AllMessageNames returns each of the current message names.
func AllMessageNames(folderPaths *FolderPaths) (names []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.messageNames: %w", err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(folderPaths.SharedMessage); err != nil {
		if os.IsNotExist(err) {
			// The folder has not been created yet.
			err = nil
			return
		}
	}
	lDirEntrys := len(dirEntrys)
	if lDirEntrys == 0 {
		// No files yet.
		return
	}
	names = make([]string, 0, len(dirEntrys))
	lExt := len(goFileExt)
	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			// Ignore directories. (shouldn't be any)
			// Only want .go files.
			continue
		}
		fileName := dirEntry.Name()
		ext := filepath.Ext(fileName)
		if ext != goFileExt {
			continue
		}
		l := len(fileName) - lExt
		messageName := Cap(fileName[:l])
		if messageName == InitMessageName {
			continue
		}
		if isValid, _ := validateMessageName(messageName); !isValid {
			continue
		}
		names = append(names, messageName)
	}
	return
}

// ValidateNewMessageName returns an error if the message name is not valid.
func ValidateNewMessageName(
	messageName string,
	folderPaths *FolderPaths,
) (isValid bool, userMessage string, err error) {

	lc := strings.ToLower(messageName)

	switch {
	case lc == strings.ToLower(chansMessageName):
		userMessage = fmt.Sprintf(chansMessageNameInvalidF, messageName)
		return
	case lc == strings.ToLower(InitMessageName):
		userMessage = fmt.Sprintf(initMessageNameInvalidF, messageName)
		return
	default:
		if isValid, userMessage = validateMessageName(messageName); !isValid {
			return
		}
	}

	var messageNames []string
	if messageNames, err = AllMessageNames(folderPaths); err != nil {
		return
	}
	for _, name := range messageNames {
		if strings.ToLower(name) == lc {
			isValid = false
			userMessage = fmt.Sprintf("The message name %q is too smilar to the message name %q.", messageName, name)
			return
		}
	}
	isValid = true
	return
}

// ValidateCurrentMessageName returns an error if the message name is not valid.
func ValidateCurrentMessageName(
	messageName string,
	folderPaths *FolderPaths,
) (isValid bool, userMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ValidateCurrentMessageName: %w", err)
		}
	}()

	switch {
	case messageName == InitMessageName:
		userMessage = initMessageNameInvalid
		return
	case messageName == chansMessageName:
		userMessage = chansMessageNameInvalid
		return
	default:
		if isValid, userMessage = validateMessageName(messageName); !isValid {
			return
		}
	}

	var messageNames []string
	if messageNames, err = UserMessageNames(folderPaths); err != nil {
		return
	}
	for _, name := range messageNames {
		if name == messageName {
			return
		}
	}
	userMessage = fmt.Sprintf("The message name %q is not being used.", messageName)
	isValid = false
	return
}
