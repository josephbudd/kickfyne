package shared

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/shared/message"
	"github.com/josephbudd/kickfyne/source/shared/meta"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	FolderName = "shared"
)

// CreateFramework creates the shared/ files.
func CreateFramework(
	appName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shared.CreateFramework: %w", err)
		}
	}()

	// shared/message/
	if err = message.CreateFramework(folderPaths); err != nil {
		return
	}

	// shared/meta/
	if err = meta.CreateFramework(importPrefix, folderPaths); err != nil {
		return
	}

	return
}
