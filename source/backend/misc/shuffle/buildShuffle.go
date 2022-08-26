package shuffle

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	fileName = "shuffle.go"
)

func Build(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shuffle.Build: %w", err)
		}
	}()

	// misc/shuffle/shuffle.go
	oPath := filepath.Join(folderPaths.BackendMiscShuffle, fileName)
	if err = utils.ProcessTemplate(fileName, oPath, shuffleTemplate, nil); err != nil {
		return
	}

	return
}
