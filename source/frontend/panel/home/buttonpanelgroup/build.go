package buttonpanelgroup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/source/utils"
)

// Build builds a button's panel group.
func Build(
	parentGroupName string,
	homeButtonName string,
	defaultPanelName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelgroup.Build: %w", err)
		}
	}()

	// folderPaths.FrontendPanelHome

	panelGroupPackageName, panelGroupGroupName, panelGroupPackagePath := Names(homeButtonName, parentGroupName, folderPaths.FrontendPanelHome)
	if err = os.Mkdir(panelGroupPackagePath, utils.DMode); err != nil {
		return
	}

	var fPath string
	var data interface{}

	// home/<home button name>PanelGroup/panelGroup.go
	fPath = filepath.Join(panelGroupPackagePath, groupFileName)
	data = groupTemplateData{
		HomeButtonName:   homeButtonName,
		GroupName:        panelGroupGroupName,
		DefaultPanelName: defaultPanelName,
		PackageName:      panelGroupPackageName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(groupFileName, fPath, groupTemplate, data); err != nil {
		return
	}

	// home/<home button name>PanelGroup/messageHandler.go
	fPath = filepath.Join(panelGroupPackagePath, messengerFileName)
	data = messengerTemplateData{
		PackageName:  panelGroupPackageName,
		ImportPrefix: importPrefix,
		GroupName:    panelGroupGroupName,
	}
	if err = utils.ProcessTemplate(messengerFileName, fPath, messengerTemplate, data); err != nil {
		return
	}

	return
}

func BuildPanel(
	panelName string,
	panelDescription string,
	panelHeading string,
	parentGroupName string,
	homeButtonName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelgroup.BuildPanel: %w", err)
		}
	}()

	panelGroupPackageName, _, panelGroupPackagePath := Names(homeButtonName, parentGroupName, folderPaths.FrontendPanelHome)

	// home/<home button name>PanelGroup/<panel name>Panel.go
	fName := utils.PanelFileName(panelName)
	fPath := filepath.Join(panelGroupPackagePath, fName)
	messengerdata := panelTemplateData{
		PanelName:        panelName,
		PanelDescription: panelDescription,
		PanelHeading:     panelHeading,
		PackageName:      panelGroupPackageName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(fName, fPath, panelTemplate, messengerdata); err != nil {
		return
	}

	return
}

func Names(
	homeButtonName string,
	parentGroupName string,
	parentFolderPath string,
) (packageName, groupName, path string) {
	folder := utils.ButtonPanelGroupFolderName(homeButtonName)
	packageName = strings.ToLower(folder)
	groupName = utils.GroupName(parentGroupName, folder)
	path = filepath.Join(parentFolderPath, folder)
	return
}
