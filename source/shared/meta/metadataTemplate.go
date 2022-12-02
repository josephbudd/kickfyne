package meta

type metadataTemplateData struct {
	ImportPrefix string
}

const (
	metadataFileName = "meta.go"

	metadataTemplate = `package meta

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"

	toml "github.com/pelletier/go-toml"
)

type FrontEndData struct {
	Landing string
}

type AppData struct {
	Details  fyne.AppMetadata
	FrontEnd FrontEndData
}

var appData = AppData{}
var loadErr error

// Data returns the application's meta data.
func Data() (data AppData, err error) {
	if err = loadErr; err != nil {
		return
	}
	if data = appData; len(data.FrontEnd.Landing) > 0 {
		return
	}
	loadMetaData()
	err = loadErr
	data = appData
	return
}

func loadMetaData() {

	var err error
	defer func() {
		if err != nil {
			err = fmt.Errorf("meta.loadMetaData: %w", err)
			loadErr = err
		}
	}()

	// Meta data.
	pwd := os.Getenv("PWD")
	metaDataPath := filepath.Join(pwd, "FyneApp.toml")
	metaDataURI := storage.NewFileURI(metaDataPath)
	// Open.
	var rc fyne.URIReadCloser
	if rc, err = storage.Reader(metaDataURI); err != nil {
		return
	}
	decoder := toml.NewDecoder(rc)
	err = decoder.Decode(&appData)
}

`
)
