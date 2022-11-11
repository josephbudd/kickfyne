package folder

import "github.com/josephbudd/kickfyne/source/utils"

const (
	folderName = "folder"
	fileName   = "folder.go"
)

type templateData struct {
	AppName string
	Funcs   utils.Funcs
}

var template = `// package paths manages file paths.
package folder

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

var (
	appDataPath     string
	appDataURI      fyne.URI
	createFolderErr error
)

// Path returns the app's data path.
func Path() (path string, err error) {
	if err = createFolderErr; err != nil {
		return
	}
	if path = appDataPath; len(path) > 0 {
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("folder.Path: %w", err)
			createFolderErr = err
		}
	}()

	// Build appDataPath.
	var homePath string
	switch runtime.GOOS {
	case "darwin":
		homePath = os.Getenv("HOME")
	case "windows":
		homePath = os.Getenv("LOCALAPPDATA")
	default:
		homePath = os.Getenv("HOME")
	}
	appDataPath = filepath.Join(homePath, ".{{ call .Funcs.LowerCase .AppName }}")
	appDataURI := storage.NewFileURI(appDataPath)
	if err = storage.CreateListable(appDataURI); !errors.Is(err, fs.ErrExist) {
		// The error does not indicate that the folder already exists.
		// It indicates some other error.
		return
	}
	err = nil
	path = appDataPath
	return
}

// FileURI returns the URI of a file in the app's data folder.
func FileURI(filename string) (uri fyne.URI, err error) {
	if err = createFolderErr; err != nil {
		return
	}
	if uri = appDataURI; uri != nil {
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("folder.FileURI: %w", err)
		}
	}()

	var folderPath string
	if folderPath, err = Path(); err != nil {
		return
	}
	filePath := filepath.Join(folderPath, filename)
	appDataURI = storage.NewFileURI(filePath)
	uri = appDataURI
	return
}

`
