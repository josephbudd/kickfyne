package state

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("state.Build: %w", err)
		}
	}()

	var oPath string

	// state/backend.go
	oPath = filepath.Join(folderPaths.SharedState, backendFileName)
	if err = utils.ProcessTemplate(backendFileName, oPath, backendTemplate, nil); err != nil {
		return
	}

	// state/frontend.go
	oPath = filepath.Join(folderPaths.SharedState, frontendFileName)
	if err = utils.ProcessTemplate(frontendFileName, oPath, frontendTemplate, nil); err != nil {
		return
	}

	// state/message.go
	oPath = filepath.Join(folderPaths.SharedState, messsageFileName)
	if err = utils.ProcessTemplate(messsageFileName, oPath, messsageTemplate, nil); err != nil {
		return
	}

	// state/messenger.go
	oPath = filepath.Join(folderPaths.SharedState, messsengerFileName)
	if err = utils.ProcessTemplate(messsengerFileName, oPath, messsengerTemplate, nil); err != nil {
		return
	}

	// state/state.go
	oPath = filepath.Join(folderPaths.SharedState, stateFileName)
	data := stateTemplateData{
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(stateFileName, oPath, stateTemplate, data); err != nil {
		return
	}

	return
}
