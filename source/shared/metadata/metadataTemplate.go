package metadata

type metadataTemplateData struct {
	ImportPrefix string
}

const (
	metadataFileName = "metadata.go"

	metadataTemplate = `package metadata

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"

	"{{ .ImportPrefix }}/shared/paths"

	toml "github.com/pelletier/go-toml"
)

type FyneAppMetaData struct {
	Details fyne.AppMetadata
}

var appData = FyneAppMetaData{}
var loadErr error
var loaded bool

func AppMetaData() (details fyne.AppMetadata, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("metadata.AppMetaData: %w", err)
		}
	}()

	if !loaded {
		loadMetaData()
	}
	if err = loadErr; err != nil {
		return
	}
	details = fyne.AppMetadata(appData.Details)
	return
}

func loadMetaData() {
	loaded = true
	// Open.
	var rc fyne.URIReadCloser
	if rc, loadErr = storage.Reader(paths.MetaDataURI()); loadErr != nil {
		return
	}
	decoder := toml.NewDecoder(rc)
	loadErr = decoder.Decode(&appData)
}

`
)
