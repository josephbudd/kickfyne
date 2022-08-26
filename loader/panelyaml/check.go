package panelyaml

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Check(panelYAML YAML, folderPaths *utils.FolderPaths) (isOK bool, failureMessage string, err error) {

	var msg string
	msgs := make([]string, 0, 4)
	defer func() {
		if err != nil {
			err = fmt.Errorf("panelyaml.Check: %w", err)
			return
		}
		if len(msgs) > 0 {
			isOK = false
			failureMessage = strings.Join(msgs, "\n")
		}
	}()

	if len(panelYAML.ButtonName) == 0 {
		msg = "The button name is missing."
		msgs = append(msgs, msg)
	}
	// Panel Name.
	if len(panelYAML.Panel.Name) == 0 {
		msg = "The panel has no name."
		msgs = append(msgs, msg)
	}
	if len(panelYAML.Panel.Heading) == 0 {
		msg = "The panel has no heading."
		msgs = append(msgs, msg)
	}
	if len(panelYAML.Panel.Description) == 0 {
		msg = "The panel has no description."
		msgs = append(msgs, msg)
	}

	if isOK, msg, err = utils.ValidateCurrentButtonName(panelYAML.ButtonName, folderPaths); !isOK || err != nil {
		if !isOK {
			msgs = append(msgs, msg)
		}
		return
	}
	var haveTab bool
	if len(panelYAML.TabName) > 0 {
		if isOK, msg, err = utils.ValidateCurrentTabName(panelYAML.ButtonName, panelYAML.TabName, folderPaths); !isOK || err != nil {
			if !isOK {
				msgs = append(msgs, msg)
			}
			return
		}
		haveTab = true
	}
	if haveTab {
		// Adding the panel to the tab.
		isOK, msg, err = utils.ValidateNewTabPanelName(
			panelYAML.ButtonName,
			panelYAML.TabName,
			panelYAML.Panel.Name,
			folderPaths,
		)
	} else {
		// Adding the panel to the button.
		isOK, msg, err = utils.ValidateNewButtonPanelName(
			panelYAML.ButtonName,
			panelYAML.Panel.Name,
			folderPaths,
		)
	}
	if !isOK {
		msgs = append(msgs, msg)
	}
	return
}
