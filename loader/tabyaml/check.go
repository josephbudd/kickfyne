package tabyaml

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Check(tabYAML YAML, folderPaths *utils.FolderPaths) (isOK bool, userMessage string, err error) {

	var msg string
	msgs := make([]string, 0, 4)
	defer func() {
		if isOK = len(msgs) == 0; !isOK {
			userMessage = strings.Join(msgs, "\n")
		}
	}()

	buttonName := tabYAML.ButtonName
	// Button Name.
	if len(buttonName) == 0 {
		msgs = append(msgs, userMessage)
	}
	var isValid bool
	if isValid, msg, err = utils.ValidateCurrentButtonName(buttonName, folderPaths); !isValid || err != nil {
		msgs = append(msgs, msg)
		return
	}
	// The button name is valid.
	// Tab
	tab := tabYAML.Tab
	if isValid, msg, err = utils.ValidateCurrentTabName(buttonName, tab.Name, folderPaths); !isValid || err != nil {
		if !isValid {
			msgs = append(msgs, msg)
		}
		return
	}
	// Panels.
	lp := len(tab.Panels)
	switch {
	case lp == 0:
		msgs = append(msgs, "The tab has no panels.")
	case lp > 0:
		// A panel group.
		if msg = checkTabPanels(tab); len(msg) > 0 {
			msgs = append(msgs, msg)
		}
	}
	return
}

func checkTabPanels(tab loader.Tab) (errorMessage string) {

	lp := len(tab.Panels)
	var msgs = make([]string, 0, lp)
	var msg string
	defer func() {
		if len(msgs) > 0 {
			errorMessage = strings.Join(msgs, "\n")
		}
	}()

	panelNames := make([]string, lp)
	panelHeadings := make([]string, lp)
	for i, panel := range tab.Panels {
		panelNames[i] = panel.Name
		panelHeadings[i] = panel.Heading
	}
	for i, panel := range tab.Panels {
		if len(panel.Name) > 0 {
			for j, panelName := range panelNames {
				if j == i {
					continue
				}
				if panelName == panel.Name {
					msg = fmt.Sprintf("The tab has multiple panels with the Name %q.", panel.Name)
					msgs = append(msgs, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The tab: Panels[%d] has no Name.", i)
			msgs = append(msgs, msg)
		}
		if len(panel.Heading) > 0 {
			for j, panelHeading := range panelHeadings {
				if j == i {
					continue
				}
				if panelHeading == panel.Heading {
					msg = fmt.Sprintf("The button has multiple Panels with the Heading %q.", panel.Heading)
					msgs = append(msgs, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The tab: Panels[%d] has no Heading.", i)
			msgs = append(msgs, msg)
		}
	}
	return
}
