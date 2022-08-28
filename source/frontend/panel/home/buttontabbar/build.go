/*
	1. create group
	groupName is package name
	group.go is the group file. // template needs each panel name.
	<panelName>.go
*/
// package buttontabbar creates a tab bar panel group for a home button
package buttontabbar

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickfyne/source/utils"
)

// Build builds the home button's tab bar folder.
func Build(
	parentGroupName string,
	homeButtonName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("buttontabbar.Build: %w", err)
		}
	}()

	tabbarPackageName, tabbarGroupName, tabbarPackagePath := Names(homeButtonName, parentGroupName, folderPaths)
	if err = os.Mkdir(tabbarPackagePath, utils.DMode); err != nil {
		return
	}

	var oPath string
	var data interface{}

	// home/<home button name>TabBar/tabBar.go.
	data = tabBarTemplateData{
		HomeButtonName: homeButtonName,
		ImportPrefix:   importPrefix,
		PackageName:    tabbarPackageName,
		Funcs:          utils.GetFuncs(),
	}
	oPath = filepath.Join(tabbarPackagePath, tabBarFileName)
	if err = utils.ProcessTemplate(tabBarFileName, oPath, tabBarTemplate, data); err != nil {
		return
	}

	// home/<home button name>TabBar/messageHandler.go.
	oPath = filepath.Join(tabbarPackagePath, messengerFileName)
	data = messengerTemplateData{
		HomeButtonName: homeButtonName,
		ImportPrefix:   importPrefix,
		PackageName:    tabbarPackageName,
		Funcs:          utils.GetFuncs(),
		GroupName:      tabbarGroupName,
	}
	if err = utils.ProcessTemplate(messengerFileName, oPath, messengerTemplate, data); err != nil {
		return
	}

	return
}

// BuildTab builds a single tab for a tabbar.
func BuildTab(
	parentGroupName string,
	homeButtonName string,
	tabName string,
	tabLabel string,
	tabIndex int,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (groupName, tabbarPackagePath string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("buttontabbar.BuildTab: %w", err)
		}
	}()

	var tabbarPackageName string
	tabbarPackageName, groupName, tabbarPackagePath = Names(homeButtonName, parentGroupName, folderPaths)

	var oPath string
	var data interface{}
	var fileName string

	// home/<home button name>TabBar/<tab name>Tab.go.
	data = tabTemplateData{
		TabName:                tabName,
		TabLabel:               tabLabel,
		TabIndex:               tabIndex,
		PackageName:            tabbarPackageName,
		HomeButtonName:         homeButtonName,
		PanelGroupFolderImport: utils.ButtonTabBarTabPanelGroupFolderImport(importPrefix, homeButtonName, tabName),
		ImportPrefix:           importPrefix,
		Funcs:                  utils.GetFuncs(),
	}
	fileName = tabFileName(tabName)
	oPath = filepath.Join(tabbarPackagePath, fileName)
	if err = utils.ProcessTemplate(fileName, oPath, tabTemplate, data); err != nil {
		return
	}

	return
}

func Names(
	homeButtonName string,
	parentGroupName string,
	folderPaths *utils.FolderPaths,
) (packageName, groupName, path string) {
	folder := utils.ButtonTabBarFolderName(homeButtonName)
	packageName = strings.ToLower(folder)
	groupName = utils.GroupName(parentGroupName, folder)
	path = filepath.Join(folderPaths.FrontendPanelHome, folder)
	return
}
