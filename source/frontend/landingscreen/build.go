package landingscreen

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func BuildLanding(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.BuildLanding: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// frontend/landing.go
	var metaData utils.FyneAppMetaData
	if metaData, err = utils.ReadMetaData(folderPaths); err != nil {
		return
	}
	oPath = filepath.Join(folderPaths.FrontendLanding, landingFileName)
	data = landingTemplateData{
		ImportPrefix:      importPrefix,
		LandingScreenName: metaData.FrontEnd.Landing,
	}
	err = utils.ProcessTemplate(landingFileName, oPath, landingTemplate, data)
	return
}
