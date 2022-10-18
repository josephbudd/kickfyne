package screens

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

// BuildPackageWithoutPanels builds screens package.
func BuildPackageWithoutPanels(
	packageName string,
	packageDoc string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.Build: %w", err)
		}
	}()

	// package folder
	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	if err = os.Mkdir(packagePath, utils.DMode); err != nil {
		return
	}

	var fPath string
	var data interface{}

	// gui/screens/«screen-package-name»/doc.go
	fPath = filepath.Join(packagePath, utils.DocFileName)
	data = docTemplateData{
		PackageName: packageName,
		PackageDoc:  packageDoc,
		Funcs:       utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.DocFileName, fPath, docTemplate, data); err != nil {
		return
	}

	// gui/screens/«screen-package-name»/screen.go
	fPath = filepath.Join(packagePath, utils.ScreenFileName)
	data = screenTemplateData{
		PackageName:  packageName,
		PanelNames:   []string{},
		ImportPrefix: importPrefix,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.ScreenFileName, fPath, screenTemplate, data); err != nil {
		return
	}

	// gui/screens/«screen-package-name»/messageHandler.go
	fPath = filepath.Join(packagePath, messengerFileName)
	data = messengerTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(messengerFileName, fPath, messengerTemplate, data); err != nil {
		return
	}

	return
}

// RemovePanelFile removes the panel file and then updates the screen.go file is required.
func RemovePanelFile(
	packageName string,
	panelName string,
	hasAccordionPanel bool,
	hasAppTabsPanel bool,
	hasDocTabsPanel bool,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.RemovePanelFile: %w", err)
		}
	}()

	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	fName := utils.PanelFileName(panelName)
	fPath := filepath.Join(packagePath, fName)
	if err = os.Remove(fPath); err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			return
		}
	}

	// Rebuild screen.go with all panel names.
	var panelNames []string
	if panelNames, err = utils.PanelNames(packagePath); err != nil {
		return
	}
	var defaultPanelName string
	switch {
	case hasAccordionPanel:
		defaultPanelName = utils.AccordionPanelName
	case hasAppTabsPanel:
		defaultPanelName = utils.AppTabsPanelName
	case hasDocTabsPanel:
		defaultPanelName = utils.DocTabsPanelName
	default:
		if len(panelNames) > 0 {
			defaultPanelName = panelNames[0]
		}
	}
	// gui/screens/«screen-package-name»/screen.go
	fPath = filepath.Join(packagePath, utils.ScreenFileName)
	screentemplatedata := screenTemplateData{
		PackageName:      packageName,
		PanelNames:       panelNames,
		DefaultPanelName: defaultPanelName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.ScreenFileName, fPath, screenTemplate, screentemplatedata); err != nil {
		return
	}
	return
}

func BuildAccordionPanelFile(
	packageName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.BuildAccordionPanelFile: %w", err)
		}
	}()

	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	fPath := filepath.Join(packagePath, utils.AccordionPanelFileName)
	paneltemplatedata := accordionPanelTemplateData{
		PackageName:  packageName,
		PanelName:    utils.AccordionPanelName,
		ImportPrefix: importPrefix,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.AccordionPanelFileName, fPath, accordionPanelTemplate, paneltemplatedata); err != nil {
		return
	}
	// Rebuild screen.go with all panel names.
	err = reBuildScreenFile(
		packageName,
		packagePath,
		utils.AccordionPanelName,
		importPrefix,
		folderPaths,
	)
	return
}

func BuildAppTabsPanelFile(
	packageName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.BuildAppTabsPanelFile: %w", err)
		}
	}()

	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	fPath := filepath.Join(packagePath, utils.AppTabsPanelFileName)
	paneltemplatedata := appTabsPanelData{
		PackageName:  packageName,
		PanelName:    utils.AppTabsPanelName,
		ImportPrefix: importPrefix,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.AppTabsPanelFileName, fPath, appTabsPanelTemplate, paneltemplatedata); err != nil {
		return
	}
	// Rebuild screen.go with all panel names.
	err = reBuildScreenFile(
		packageName,
		packagePath,
		utils.AppTabsPanelName,
		importPrefix,
		folderPaths,
	)
	return
}

func BuildDocTabsPanelFile(
	packageName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.BuildDocTabsPanelFile: %w", err)
		}
	}()

	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	fPath := filepath.Join(packagePath, utils.DocTabsPanelFileName)
	paneltemplatedata := doctabsPanelData{
		PackageName:  packageName,
		PanelName:    utils.DocTabsPanelName,
		ImportPrefix: importPrefix,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(utils.DocTabsPanelFileName, fPath, doctabsPanelTemplate, paneltemplatedata); err != nil {
		return
	}
	// Rebuild screen.go with all panel names.
	err = reBuildScreenFile(
		packageName,
		packagePath,
		utils.DocTabsPanelName,
		importPrefix,
		folderPaths,
	)
	return
}

func BuildPanelFile(
	packageName string,
	panelName string,
	hasAccordionPanel bool,
	hasAppTabsPanel bool,
	hasDocTabsPanel bool,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	// gui/screens/«screen-package-name»/«panel name»Panel.go
	fName := utils.PanelFileName(panelName)
	packagePath := filepath.Join(folderPaths.FrontendGUIScreens, packageName)
	fPath := filepath.Join(packagePath, fName)
	paneltemplatedata := panelTemplateData{
		PanelName:    panelName,
		PackageName:  packageName,
		AddShowFunc:  !(hasAccordionPanel || hasAppTabsPanel || hasDocTabsPanel),
		ImportPrefix: importPrefix,
		Funcs:        utils.GetFuncs(),
	}
	if err = utils.ProcessTemplate(fName, fPath, panelTemplate, paneltemplatedata); err != nil {
		return
	}
	var defaultPanelName string
	switch {
	case hasAccordionPanel:
		defaultPanelName = utils.AccordionPanelName
	case hasAppTabsPanel:
		defaultPanelName = utils.AppTabsPanelName
	case hasDocTabsPanel:
		defaultPanelName = utils.DocTabsPanelName
	}
	// Rebuild screen.go with all panel names.
	// This may change the screens default content to another panel.
	err = reBuildScreenFile(
		packageName,
		packagePath,
		defaultPanelName,
		importPrefix,
		folderPaths,
	)

	return
}

// reBuildScreenFile rebuilds the screen.go file making sure each panel's func Init is called.
func reBuildScreenFile(
	packageName string,
	packagePath string,
	defaultPanelName string,
	importPrefix string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("screens.reBuildScreenFile: %w", err)
		}
	}()

	var panelNames []string
	if panelNames, err = utils.PanelNames(packagePath); err != nil {
		return
	}
	if len(defaultPanelName) == 0 {
		if len(panelNames) > 0 {
			defaultPanelName = panelNames[0]
		}
	}
	// gui/screens/«screen-package-name»/screen.go
	fPath := filepath.Join(packagePath, utils.ScreenFileName)
	data := screenTemplateData{
		PackageName:      packageName,
		PanelNames:       panelNames,
		DefaultPanelName: defaultPanelName,
		ImportPrefix:     importPrefix,
		Funcs:            utils.GetFuncs(),
	}
	err = utils.ProcessTemplate(utils.ScreenFileName, fPath, screenTemplate, data)
	return
}
