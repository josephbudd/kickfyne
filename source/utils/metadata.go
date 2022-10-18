package utils

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"github.com/BurntSushi/toml"
)

type FrontEndMetaData struct {
	Landing string
}

type FyneAppMetaData struct {
	Details  fyne.AppMetadata
	FrontEnd FrontEndMetaData
}

var loadedAppData *FyneAppMetaData
var loadErr error

func ReadMetaData(folderPaths *FolderPaths) (data FyneAppMetaData, err error) {

	if loadedAppData != nil || loadErr != nil {
		data = *loadedAppData
		err = loadErr
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("metadata.AppMetaData: %w", err)
		}
	}()

	data, err = loadMetaData(folderPaths)
	loadedAppData = &data
	loadErr = err
	return
}

func loadMetaData(folderPaths *FolderPaths) (data FyneAppMetaData, err error) {

	var file *os.File
	defer func() {
		if cErr := file.Close(); cErr != nil {
			if err == nil {
				err = cErr
			}
		}
		if err != nil {
			err = fmt.Errorf("fyneapptoml.loadMetaData: %w", err)
			return
		}
	}()

	path := FyneAppTOMLFilePath(folderPaths)
	if file, err = os.Open(path); err != nil {
		return
	}
	decoder := toml.NewDecoder(file)
	_, err = decoder.Decode(&data)
	return
}
