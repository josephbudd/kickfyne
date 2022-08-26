package widget

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/frontend/widget/backpanel"
	"github.com/josephbudd/kickfyne/source/frontend/widget/safebutton"
	"github.com/josephbudd/kickfyne/source/frontend/widget/selection"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("widget.Build: %w", err)
		}
	}()

	if err = safebutton.Build(folderPaths); err != nil {
		return
	}
	if err = selection.Build(folderPaths); err != nil {
		return
	}
	if err = backpanel.Build(folderPaths); err != nil {
		return
	}

	return
}
