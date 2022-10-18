package paths

import "github.com/josephbudd/kickfyne/source/utils"

const (
	folderName = "paths"
	fileName   = "paths.go"
)

type templateData struct {
	AppName string
	Funcs   utils.Funcs
}

var template = `// package paths manages file paths.
// The environment variable USETESTPATH signals that the test path is used not the normal application path.
package paths

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
	appStoresPath   string
	shareImagesPath string
	metaDataPath    string
)

func Init() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("paths.Init: %w", err)
		}
	}()

	var path string
	switch runtime.GOOS {
	case "darwin":
		path = os.Getenv("HOME")
	case "windows":
		path = os.Getenv("LOCALAPPDATA")
	default:
		path = os.Getenv("HOME")
	}
	if len(os.Getenv("USETESTPATH")) > 0 {
		appDataPath = filepath.Join(path, ".{{ call .Funcs.LowerCase .AppName }}_test")
	} else {
		appDataPath = filepath.Join(path, ".{{ call .Funcs.LowerCase .AppName }}")
	}

	// The app's folder.
	appDataURI := storage.NewFileURI(appDataPath)
	if err = storage.CreateListable(appDataURI); err != nil {
		if !errors.Is(err, fs.ErrExist) {
			// The error does not indicate that the folder already exists.
			// It indicates some other error.
			return
		}
	}

	// The stores folder in the app's folder.
	appStoresPath = filepath.Join(appDataPath, "stores")
	appStoresURI := storage.NewFileURI(appStoresPath)
	if err = storage.CreateListable(appStoresURI); err != nil {
		if !errors.Is(err, fs.ErrExist) {
			// The error does not indicate that the folder already exists.
			// It indicates some other error.
			return
		}
		err = nil
	}

	// Images folder.
	pwd := os.Getenv("PWD")
	shareImagesPath = filepath.Join(pwd, "images")

	// Meta data.
	metaDataPath = filepath.Join(pwd, "FyneApp.toml")
	return
}

// StoreURI returns the URI of the folder where records are stored by default.
func StoreURI(filename string) (storeURI fyne.URI) {
	path := filepath.Join(appStoresPath, filename)
	storeURI = storage.NewFileURI(path)
	return
}

// ImageURI returns the URI of the images folder.
func ImageURI(filename string) (imageURI fyne.URI) {
	path := filepath.Join(shareImagesPath, filename)
	imageURI = storage.NewFileURI(path)
	return
}

// MetaDataURI returns the URI of fyne's application meta data file.
func MetaDataURI() (metaDataURI fyne.URI) {
	metaDataURI = storage.NewFileURI(metaDataPath)
	return
}

`
