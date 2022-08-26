package home

import (
	"fmt"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func Build(
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home.Build: %w", err)
		}
	}()

	var oPath string
	var data interface{}

	// panel/home/buttonPad.go
	oPath = filepath.Join(folderPaths.FrontendPanelHome, buttonPadFileName)
	data = homeGroupTemplateData{
		ImportPrefix: importPrefix,
		//EffectImports: ,
	}
	if err = utils.ProcessTemplate(buttonPadFileName, oPath, buttonPadTemplate, data); err != nil {
		return
	}

	return
}

func BuildHomeTabBarButton(
	importPrefix string,
	folderPaths *utils.FolderPaths,
	buttonName string,
	buttonLabel string,
	buttonIndex int,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home.BuildHomeTabBarButton: %w", err)
		}
	}()

	fName := utils.HomeButtonFileName(buttonName)
	oPath := filepath.Join(folderPaths.FrontendPanelHome, fName)
	data := buttonTemplateData{
		ImportPrefix:            importPrefix,
		ButtonName:              buttonName,
		ButtonLabel:             buttonLabel,
		ButtonIndex:             buttonIndex,
		ButtonGroupFolderImport: utils.ButtonTabBarFolderImport(importPrefix, buttonName),
		Funcs:                   utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(fName, oPath, buttonTemplate, data)
	return
}

func BuildHomePanelGroupButton(
	importPrefix string,
	folderPaths *utils.FolderPaths,
	buttonName string,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home.BuildHomePanelGroupButton: %w", err)
		}
	}()

	fName := utils.HomeButtonFileName(buttonName)
	oPath := filepath.Join(folderPaths.FrontendPanelHome, fName)
	data := buttonTemplateData{
		ImportPrefix:            importPrefix,
		ButtonName:              buttonName,
		ButtonGroupFolderImport: utils.ButtonPanelGroupFolderImport(importPrefix, buttonName),
		Funcs:                   utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(fName, oPath, buttonTemplate, data)
	return
}
