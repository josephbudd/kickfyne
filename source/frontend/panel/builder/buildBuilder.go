package builder

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// Build builds the builer folder and files.
// This is where the frontend, buttons, tabbars, tabs, panelGroups, panels are registered.
func Build(
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("builder.Build: %w", err)
		}
	}()

	var oPath string

	// builder/builder.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, builderFileName)
	if err = utils.ProcessTemplate(builderFileName, oPath, builderTemplate, nil); err != nil {
		return
	}

	// builder/button.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, buttonFileName)
	if err = utils.ProcessTemplate(buttonFileName, oPath, buttonTemplate, nil); err != nil {
		return
	}

	// builder/buttonPanelGroup.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, buttonPanelGroupFileName)
	if err = utils.ProcessTemplate(buttonPanelGroupFileName, oPath, buttonPanelGroupTemplate, nil); err != nil {
		return
	}

	// builder/buttonPanelGroup.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, buttonPanelGroupFileName)
	if err = utils.ProcessTemplate(buttonFileName, oPath, buttonPanelGroupTemplate, nil); err != nil {
		return
	}

	// builder/buttonPanelGroupPanel.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, buttonPanelGroupPanelFileName)
	if err = utils.ProcessTemplate(buttonPanelGroupPanelFileName, oPath, buttonPanelGroupPanelTemplate, nil); err != nil {
		return
	}

	// builder/tab.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, tabFileName)
	if err = utils.ProcessTemplate(tabFileName, oPath, tabTemplate, nil); err != nil {
		return
	}

	// builder/tabBar.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, tabBarFileName)
	if err = utils.ProcessTemplate(tabBarFileName, oPath, tabBarTemplate, nil); err != nil {
		return
	}

	// builder/tabPanelGroup.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, tabPanelGroupFileName)
	if err = utils.ProcessTemplate(tabPanelGroupFileName, oPath, tabPanelGroupTemplate, nil); err != nil {
		return
	}

	// builder/tabPanelGroupPanel.go
	oPath = filepath.Join(folderPaths.FrontendPanelBuilder, tabPanelGroupPanelFileName)
	if err = utils.ProcessTemplate(tabPanelGroupPanelFileName, oPath, tabPanelGroupPanelTemplate, nil); err != nil {
		return
	}

	return
}
