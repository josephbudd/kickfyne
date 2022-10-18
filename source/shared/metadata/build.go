package metadata

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// CreateFramework creates the shared/ files.
func CreateFramework(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("metadata.CreateFramework: %w", err)
		}
	}()

	oPath := filepath.Join(folderPaths.SharedMetaData, metadataFileName)
	data := metadataTemplateData{
		ImportPrefix: importPrefix,
	}
	err = utils.ProcessTemplate(metadataFileName, oPath, metadataTemplate, data)
	return
}

// AppMetaDataFilePath returns the path to the app's meta data file.
func AppMetaDataFilePath(folderPaths *utils.FolderPaths) (path string) {
	path = filepath.Join(folderPaths.SharedMetaData, metadataFileName)
	return
}
