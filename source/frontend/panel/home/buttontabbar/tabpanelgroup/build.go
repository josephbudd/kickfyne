package tabpanelgroup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/source/utils"
)

// Build builds a tab's panel group.
func Build(
	homeButtonName string,
	parentGroupName string,
	tabName string,
	defaultPanelName string,
	parentFolderPath string,
	importPrefix string,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelgroup.Build: %w", err)
		}
	}()

	panelGroupPackageName, panelGroupGroupName, panelGroupPackagePath := Names(tabName, parentGroupName, parentFolderPath)
	if err = os.Mkdir(panelGroupPackagePath, utils.DMode); err != nil {
		return
	}

	var fPath string

	// home/<home button name>TabBar/<tab name>PanelGroup/panelGroup.go
	fPath = filepath.Join(panelGroupPackagePath, groupFileName)
	groupdata := groupTemplateData{
		HomeButtonName:   homeButtonName,
		GroupName:        panelGroupGroupName,
		TabName:          tabName,
		DefaultPanelName: defaultPanelName,
		PackageName:      panelGroupPackageName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(groupFileName, fPath, groupTemplate, groupdata); err != nil {
		return
	}

	// home/<home button name>TabBar/<tab name>PanelGroup/stateHandler.go
	fPath = filepath.Join(panelGroupPackagePath, staterFileName)
	staterData := staterTemplateData{
		PackageName:  panelGroupPackageName,
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(staterFileName, fPath, staterTemplate, staterData); err != nil {
		return
	}

	// home/<home button name>TabBar/<tab name>PanelGroup/messageHandler.go
	fPath = filepath.Join(panelGroupPackagePath, messengerFileName)
	messengerdata := messengerTemplateData{
		GroupName:    panelGroupGroupName,
		PackageName:  panelGroupPackageName,
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(messengerFileName, fPath, messengerTemplate, messengerdata); err != nil {
		return
	}

	return
}

func BuildPanel(
	homeButtonName string,
	parentGroupName string,
	tabName string,
	parentFolderPath string,
	panelName string,
	panelDescription string,
	panelHeading string,
	importPrefix string,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelgroup.BuildPanel: %w", err)
		}
	}()

	panelGroupPackageName, _, panelGroupPackagePath := Names(tabName, parentGroupName, parentFolderPath)

	// home/<home button name>TabBar/<tab name>PanelGroup/<panel name>Panel.go
	fName := utils.PanelFileName(panelName)
	fPath := filepath.Join(panelGroupPackagePath, fName)
	data := panelTemplateData{
		PanelName:        panelName,
		PanelDescription: panelDescription,
		PanelHeading:     panelHeading,
		PackageName:      panelGroupPackageName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(fName, fPath, panelTemplate, data); err != nil {
		return
	}

	return
}

func Names(
	tabName string,
	parentGroupName string,
	parentFolderPath string,
) (packageName, groupName, path string) {
	folder := utils.TabPanelGroupFolderName(tabName)
	packageName = strings.ToLower(folder)
	groupName = utils.GroupName(parentGroupName, folder)
	path = filepath.Join(parentFolderPath, folder)
	return
}
