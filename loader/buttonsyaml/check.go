package buttonsyaml

import (
	"fmt"
	"log"
	"strings"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/source/utils"
)

func Check(buttonsYAML YAML, folderPaths *utils.FolderPaths) (isOK bool, failureMessage string, err error) {

	msgs := make([]string, 0, len(buttonsYAML.Buttons))
	defer func() {
		if err != nil {
			err = fmt.Errorf("buttonsyaml.Check: %w", err)
			return
		}
		if len(msgs) > 0 {
			failureMessage = strings.Join(msgs, "\n")
		}
	}()

	var i, j int
	var button loader.Button
	var name, label string
	buttonNames := make([]string, len(buttonsYAML.Buttons))
	buttonLabels := make([]string, len(buttonsYAML.Buttons))
	for i, button := range buttonsYAML.Buttons {
		buttonNames[i] = button.Name
		buttonLabels[i] = button.Label
	}
	var msg string
	// Are the button names too much alike?
	for i, buttonName := range buttonNames {
		lcButtonName := strings.ToLower(buttonName)
		for j, name = range buttonNames {
			if j == i {
				continue
			}
			if name == buttonName {
				isOK = false
				msg = fmt.Sprintf("The button name %q is used multiple times.", buttonName)
				msgs = append(msgs, msg)
			}
			if strings.ToLower(name) == lcButtonName {
				isOK = false
				msg = fmt.Sprintf("The button names %q and %q are too much alike.", buttonName, name)
				msgs = append(msgs, msg)
			}
		}
	}
	// Are the labels too much alike?
	for i, buttonLabel := range buttonLabels {
		for j, label = range buttonLabels {
			if j == i {
				continue
			}
			if label == buttonLabel {
				isOK = false
				msg = fmt.Sprintf("The button label %q is used multiple times.", buttonLabel)
				msgs = append(msgs, msg)
			}
		}
	}

	// Each button name must be new.
	if isOK, msg, err = utils.ValidateNewButtonNames(buttonNames, folderPaths); !isOK || err != nil {
		if !isOK {
			msgs = append(msgs, msg)
		}
		return
	}

	names := make([]string, len(buttonsYAML.Buttons))
	labels := make([]string, len(buttonsYAML.Buttons))
	for i, button = range buttonsYAML.Buttons {
		names[i] = button.Name
		labels[i] = button.Label
	}
	for i, button = range buttonsYAML.Buttons {

		// What to call this button.
		var buttonName string
		switch {
		case len(button.Name) > 0:
			buttonName = fmt.Sprintf("Button named %q", button.Name)
		case len(button.Label) > 0:
			buttonName = fmt.Sprintf("Button labeled %q", button.Label)
		default:
			buttonName = fmt.Sprintf("Buttons[%d]", i)
		}

		// Button Name.
		if len(button.Name) > 0 {
			for j, name := range names {
				if j == i {
					continue
				}
				if name == button.Name {
					msg = fmt.Sprintf("The button name %q is used multiple times.", button.Name)
					msgs = append(msgs, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s has no name.", buttonName)
			msgs = append(msgs, msg)
		}

		// Button Label.
		if len(button.Label) > 0 {
			for j, label := range labels {
				if j == i {
					continue
				}
				if label == button.Label {
					msg = fmt.Sprintf("The button label %q is used multiple times.", button.Label)
					msgs = append(msgs, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s has no label.", buttonName)
			msgs = append(msgs, msg)
		}

		// Panels and Tabs.
		lp := len(button.Panels)
		lt := len(button.Tabs)
		switch {
		case lp > 0 && lt > 0:
			msg = fmt.Sprintf("The %s has %d panels and %d tabs.", buttonName, lp, lt)
			msgs = append(msgs, msg)
		case lp == 0 && lt == 0:
			msg = fmt.Sprintf("The %s has %d panels and %d tabs.", buttonName, lp, lt)
			msgs = append(msgs, msg)
		case lt > 0 && lp == 0:
			// A tab bar panel.
			if isOK, msg = checkButtonTabs(buttonName, button); !isOK {
				msgs = append(msgs, msg)
			}
		case lt == 0 && lp > 0:
			// A panel group.
			if errorMessages := checkButtonPanels(buttonName, button); len(errorMessages) > 0 {
				msgs = append(msgs, errorMessages...)
			}
		}
	}
	return
}

func checkButtonPanels(buttonName string, button loader.Button) (errorMessages []string) {
	lp := len(button.Panels)
	names := make([]string, lp)
	headings := make([]string, lp)
	for i, panel := range button.Panels {
		names[i] = panel.Name
		headings[i] = panel.Heading
	}
	errorMessages = make([]string, 0, lp)
	var msg string
	for i, panel := range button.Panels {
		if len(panel.Name) > 0 {
			for j, name := range names {
				if j == i {
					continue
				}
				if name == panel.Name {
					log.Printf("i is %d, j is %d: name is %q.", i, j, name)
					msg = fmt.Sprintf("The %s has multiple panels with the Name %q.", buttonName, panel.Name)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s :Panels[%d] has no Name.", buttonName, i)
			errorMessages = append(errorMessages, msg)
		}
		if len(panel.Heading) > 0 {
			for j, heading := range headings {
				if j == i {
					continue
				}
				if heading == panel.Heading {
					msg = fmt.Sprintf("The %s has multiple Panels with the Heading %q.", buttonName, panel.Heading)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s :Panels[%d] has no Heading.", buttonName, i)
			errorMessages = append(errorMessages, msg)
		}
	}
	return
}

func checkButtonTabs(buttonName string, button loader.Button) (isOK bool, errorMessage string) {

	var errorMessages []string
	defer func() {
		if isOK = len(errorMessages) == 0; !isOK {
			errorMessage = strings.Join(errorMessages, "\n")
		}
	}()

	lt := len(button.Tabs)
	names := make([]string, lt)
	labels := make([]string, lt)
	for i, tab := range button.Tabs {
		names[i] = tab.Name
		labels[i] = tab.Label
	}
	errorMessages = make([]string, 0, lt)
	var msg string
	for i, tab := range button.Tabs {

		// What to call this tab.
		var tabName string
		switch {
		case len(tab.Name) > 0:
			tabName = fmt.Sprintf("Tab named %q", tab.Name)
		case len(tab.Label) > 0:
			tabName = fmt.Sprintf("Tab labeled %q", tab.Label)
		default:
			tabName = fmt.Sprintf("Tabs[%d]", i)
		}

		if len(tab.Name) > 0 {
			for j, name := range names {
				if j == i {
					continue
				}
				if name == tab.Name {
					msg = fmt.Sprintf("The %s has multiple Tabs with the Name %q.", buttonName, tab.Name)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s :Tabs[%d] has no Name.", buttonName, i)
			errorMessages = append(errorMessages, msg)
		}
		if len(tab.Label) > 0 {
			for j, label := range labels {
				if j == i {
					continue
				}
				if label == tab.Label {
					msg = fmt.Sprintf("The %s has multiple Tabs with the Label %q.", buttonName, tab.Name)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s: %s has no Label.", buttonName, tabName)
			errorMessages = append(errorMessages, msg)
		}
		if len(tab.Panels) == 0 {
			msg = fmt.Sprintf("The %s: %s has no Panels.", buttonName, tabName)
			errorMessages = append(errorMessages, msg)
		} else {
			msgs := checkButtonTabPanels(buttonName, tabName, tab)
			errorMessages = append(errorMessages, msgs...)
		}
	}
	return
}

func checkButtonTabPanels(buttonName, tabName string, tab loader.Tab) (errorMessages []string) {
	lp := len(tab.Panels)
	names := make([]string, lp)
	headings := make([]string, lp)
	for i, panel := range tab.Panels {
		names[i] = panel.Name
		headings[i] = panel.Heading
	}
	errorMessages = make([]string, 0, lp)
	var msg string
	for i, panel := range tab.Panels {
		if len(panel.Name) > 0 {
			for j, name := range names {
				if j == i {
					continue
				}
				if name == panel.Name {
					msg = fmt.Sprintf("The %s: %s has multiple panels with the Name %q.", buttonName, tabName, panel.Name)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s: %s: Panels[%d] has no Name.", buttonName, tabName, i)
			errorMessages = append(errorMessages, msg)
		}
		if len(panel.Heading) > 0 {
			for j, heading := range headings {
				if j == i {
					continue
				}
				if heading == panel.Heading {
					msg = fmt.Sprintf("The %s: %s has multiple panels with the Heading %q.", buttonName, tabName, panel.Heading)
					errorMessages = append(errorMessages, msg)
					break
				}
			}
		} else {
			msg = fmt.Sprintf("The %s: %s: Panels[%d] has no Heading.", buttonName, tabName, i)
			errorMessages = append(errorMessages, msg)
		}

	}
	return
}
