package frontend

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/frontend/panel/home"
	"github.com/josephbudd/kickfyne/source/utils"
)

func handleClear(dumperCh chan string, isBuilt bool, folderPaths *utils.FolderPaths) (err error) {

	var failureMessage string
	var successMessage string
	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.handleClear: %w", err)
			return
		}
		switch {
		case len(failureMessage) > 0:
			dumperCh <- "Failure:"
			dumperCh <- failureMessage
			dumperCh <- Usage()
		case len(successMessage) > 0:
			dumperCh <- "Success:"
			dumperCh <- successMessage
		}
	}()

	if !isBuilt {
		failureMessage = "The app must be initailized before the front end panels can be removed."
		return
	}
	// Remove buttons from the home button-panel.
	if err = home.ClearOut(folderPaths); err != nil {
		return
	}
	successMessage = "The home buttons along with any tab-bars and panel-groups are removed."
	return
}
