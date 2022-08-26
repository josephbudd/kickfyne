package frontend

import (
	"fmt"

	"github.com/josephbudd/kickfyne/loader"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttonpanelgroup"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar"
	"github.com/josephbudd/kickfyne/source/frontend/panel/home/buttontabbar/tabpanelgroup"
	"github.com/josephbudd/kickfyne/source/utils"
)

func addButton(loaderButton loader.Button, buttonIndex int, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.addButton: %w", err)
			return
		}
	}()

	// Read the button YAML file.
	if len(loaderButton.Tabs) > 0 {
		// Create the home button file.
		// panel/home/<button name>Button.go.
		if err = home.BuildHomeTabBarButton(importPrefix, folderPaths, loaderButton.Name, loaderButton.Label, buttonIndex); err != nil {
			return
		}
		if err = buttontabbar.Build(
			utils.HomeGroupName,
			loaderButton.Name,
			importPrefix,
			folderPaths,
		); err != nil {
			return
		}
		for tabIndex, tab := range loaderButton.Tabs {
			var tabGroupName string
			var tabFolderPath string
			if tabGroupName, tabFolderPath, err = buttontabbar.BuildTab(
				utils.HomeGroupName,
				loaderButton.Name,
				tab.Name,
				tab.Label,
				tabIndex,
				importPrefix,
				folderPaths,
			); err != nil {
				return
			}
			defaultPanel := tab.Panels[0]
			if err = tabpanelgroup.Build(
				loaderButton.Name,
				tabGroupName,
				tab.Name,
				defaultPanel.Name,
				tabFolderPath,
				importPrefix,
			); err != nil {
				return
			}
			for _, panel := range tab.Panels {
				if err = tabpanelgroup.BuildPanel(
					loaderButton.Name,
					tabGroupName,
					tab.Name,
					tabFolderPath,
					panel.Name,
					panel.Description,
					panel.Heading,
					importPrefix,
				); err != nil {
					return
				}
			}
		}
	} else {
		// Create the home button file.
		if err = home.BuildHomePanelGroupButton(importPrefix, folderPaths, loaderButton.Name); err != nil {
			return
		}
		defaultPanel := loaderButton.Panels[0]
		if err = buttonpanelgroup.Build(
			utils.HomeGroupName,
			loaderButton.Name,
			defaultPanel.Name,
			importPrefix,
			folderPaths,
		); err != nil {
			return
		}
		for _, panel := range loaderButton.Panels {
			if err = buttonpanelgroup.BuildPanel(
				panel.Name,
				panel.Description,
				panel.Heading,
				utils.HomeGroupName,
				loaderButton.Name,
				importPrefix,
				folderPaths,
			); err != nil {
				return
			}
		}
	}
	return
}
