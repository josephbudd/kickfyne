package buttonyaml

import (
	"fmt"
	"strings"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Check(buttonYAML YAML, folderPaths *utils.FolderPaths) (isOK bool, failureMessage string, err error) {

	var msg string
	msgs := make([]string, 0, 4)
	defer func() {
		if err != nil {
			err = fmt.Errorf("buttonyaml.Check: %w", err)
			return
		}
		if len(msgs) > 0 {
			isOK = false
			failureMessage = strings.Join(msgs, "\n")
		}
	}()

	button := buttonYAML.Button
	// Button Name.
	if len(button.Name) == 0 {
		msg = "The button has no name."
		msgs = append(msgs, msg)
		return
	} else {
		if isOK, msg, err = utils.ValidateNewButtonName(button.Name, folderPaths); !isOK || err != nil {
			if !isOK {
				msgs = append(msgs, msg)
			}
			return
		}
	}

	// Button Label.
	if len(button.Label) == 0 {
		msg = "The button has no label."
		msgs = append(msgs, msg)
	}

	// Panels and Tabs.
	lp := len(button.Panels)
	lt := len(button.Tabs)
	switch {
	case lp > 0 && lt > 0:
		msg = fmt.Sprintf("The button has %d panels and %d tabs.", lp, lt)
		msgs = append(msgs, msg)
	case lp == 0 && lt == 0:
		msg = fmt.Sprintf("The button has %d panels and %d tabs.", lp, lt)
		msgs = append(msgs, msg)
	case lt > 0 && lp == 0:
		// A tab bar panel.
		if errorMessages := checkButtonTabs(button); len(errorMessages) > 0 {
			msgs = append(msgs, errorMessages...)
		}
	case lt == 0 && lp > 0:
		// A panel group.
		if errorMessages := checkButtonPanels(button); len(errorMessages) > 0 {
			msgs = append(msgs, errorMessages...)
		}
	}
	return
}

func checkButtonPanels(button loader.Button) (errorMessages []string) {
	lp := len(button.Panels)
	panelNames := make([]string, lp)
	panelHeadings := make([]string, lp)
	for i, panel := range button.Panels {
		panelNames[i] = panel.Name
		panelHeadings[i] = panel.Heading
	}
	errorMessages = make([]string, 0, lp)
	var msg string
	for i, panel := range button.Panels {
		if len(panel.Name) > 0 {
			for j, panelName := range panelNames {
				if j == i {
					continue
				}
				if panelName == panel.Name {
					msg = fmt.Sprintf("The button has multiple panels with the Name %q.", panel.Name)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The button: Panels[%d] has no Name.", i)
			errorMessages = append(errorMessages, msg)
		}
		if len(panel.Heading) > 0 {
			for j, panelHeading := range panelHeadings {
				if j == i {
					continue
				}
				if panelHeading == panel.Heading {
					msg = fmt.Sprintf("The button has multiple Panels with the Heading %q.", panel.Heading)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The button: Panels[%d] has no Heading.", i)
			errorMessages = append(errorMessages, msg)
		}
	}
	return
}

func checkButtonTabs(button loader.Button) (errorMessages []string) {
	lt := len(button.Tabs)
	tabRefs := make([]string, lt)
	tabNames := make([]string, lt)
	tabLabels := make([]string, lt)
	for i, tab := range button.Tabs {
		tabNames[i] = tab.Name
		tabLabels[i] = tab.Label
		switch {
		case len(tab.Name) > 0:
			tabRefs[i] = fmt.Sprintf("Tabs[%d], named %q", i, tabNames[i])
		case len(tab.Label) > 0:
			tabRefs[i] = fmt.Sprintf("Tabs[%d], labeled %q", i, tabLabels[i])
		default:
			tabRefs[i] = fmt.Sprintf("Tabs[%d]", i)
		}
	}
	errorMessages = make([]string, 0, lt)
	var msg string

	var i, j int
	var tabName, name string
	// What to call each tab.
	// Check tab name redundancy.
	for i, tabName = range tabNames {
		if len(tabName) == 0 {
			msg = fmt.Sprintf("%s, has no name.", tabRefs[i])
			errorMessages = append(errorMessages, msg)
			continue
		}
		for j, name = range tabNames {
			if i == j {
				continue
			}
			if tabName == name {
				msg = fmt.Sprintf("The button has more than one tab named %q", tabName)
				errorMessages = append(errorMessages, msg)
			}
		}
	}
	// Check tab label reduncancy.
	var tabLabel, label string
	for i, tabLabel = range tabLabels {
		if len(tabLabel) == 0 {
			msg = fmt.Sprintf("%s, has no label.", tabRefs[i])
			errorMessages = append(errorMessages, msg)
			continue
		}
		for j, label = range tabLabels {
			if i == j {
				continue
			}
			if tabLabel == label {
				msg = fmt.Sprintf("The button has more than one tab labeled %q", tabLabel)
				errorMessages = append(errorMessages, msg)
			}
		}
	}
	// Check tab panels.
	var tab loader.Tab
	for i, tab = range button.Tabs {
		if len(tab.Panels) == 0 {
			msg = fmt.Sprintf("%s, has no Panels.", tabRefs[i])
			errorMessages = append(errorMessages, msg)
		} else {
			msgs := checkButtonTabPanels(tabName, tab)
			if len(msgs) > 0 {
				errorMessages = append(errorMessages, msgs...)
			}
		}
	}
	return
}

func checkButtonTabPanels(tabName string, tab loader.Tab) (errorMessages []string) {
	lp := len(tab.Panels)
	errorMessages = make([]string, 0, lp)
	var msg string
	panelNames := make([]string, lp)
	panelHeadings := make([]string, lp)
	panelDescriptions := make([]string, lp)
	panelRefs := make([]string, lp)
	for i, panel := range tab.Panels {
		panelNames[i] = panel.Name
		panelHeadings[i] = panel.Heading
		panelDescriptions[i] = panel.Description
		switch {
		case len(panel.Name) > 0:
			panelRefs[i] = fmt.Sprintf("Panels[%d] named %q", i, panelNames[i])
		case len(panel.Heading) > 0:
			panelRefs[i] = fmt.Sprintf("Panels[%d] labeled %q", i, panelHeadings[i])
		default:
			panelRefs[i] = fmt.Sprintf("Panels[%d]", i)
		}
	}
	for i, panelName := range panelNames {
		if len(panelName) == 0 {
			msg = fmt.Sprintf("%s, %s has no name.", tabName, panelRefs[i])
			errorMessages = append(errorMessages, msg)
			continue
		}
		for j, pName := range panelNames {
			if i == j {
				continue
			}
			if panelName == pName {
				msg = fmt.Sprintf("%s has multiple panels named %q.", tabName, pName)
				errorMessages = append(errorMessages, msg)
			}
		}
	}
	for i, panelHeading := range panelHeadings {
		if len(panelHeading) == 0 {
			msg = fmt.Sprintf("%s, %s has no heading.", tabName, panelRefs[i])
			errorMessages = append(errorMessages, msg)
			continue
		}
		for j, pHeading := range panelHeadings {
			if i == j {
				continue
			}
			if panelHeading == pHeading {
				msg = fmt.Sprintf("%s has multiple panels with the heading %q.", tabName, pHeading)
				errorMessages = append(errorMessages, msg)
			}
		}
	}
	return
}
