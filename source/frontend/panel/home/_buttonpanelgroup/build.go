package buttonpanelgroup

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/josephbudd/kickfyne/source/panelloader"
	"github.com/josephbudd/kickfyne/source/utils"
)

// Build builds a home buttton's panel group.
func Build(
	parentGroupName string,
	button panelloader.Button,
	importPrefix string,
	folderPaths *utils.FolderPaths,
	dumperCh chan string,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelgroup.Build: %w", err)
		}
	}()

	packageName := strings.ToLower(button.Name)
	groupName := utils.GroupName(parentGroupName, packageName)
	panelGroupPackagePath := filepath.Join(folderPaths.FrontendPanelHome, packageName)
	if err = os.Mkdir(panelGroupPackagePath, utils.DMode); err != nil {
		return
	}
	var sortedNames sort.StringSlice
	sortedNames = make([]string, len(button.Panels))
	for i, panel := range button.Panels {
		sortedNames[i] = panel.Name
	}
	sort.Sort(sortedNames)

	// <package>/group.go file.
	groupdata := groupTemplateData{
		GroupName:        groupName,
		PackageName:      packageName,
		ImportPrefix:     importPrefix,
		Button:           button,
		PanelNamesSorted: sortedNames,
		Funcs:            utils.GetFuncs(),
	}
	var fPath string
	var fName string
	fPath = filepath.Join(panelGroupPackagePath, groupFileName)
	if err = utils.ProcessTemplate(groupFileName, fPath, groupTemplate, groupdata); err != nil {
		return
	}

	// <package>/messageHandler.go
	fPath = filepath.Join(panelGroupPackagePath, messengerFileName)
	messengerdata := messengerTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(messengerFileName, fPath, messengerTemplate, messengerdata); err != nil {
		return
	}

	// <package>/stateHandler.go
	fPath = filepath.Join(panelGroupPackagePath, staterFileName)
	staterData := staterTemplateData{
		PackageName:  packageName,
		ImportPrefix: importPrefix,
	}
	if err = utils.ProcessTemplate(staterFileName, fPath, staterTemplate, staterData); err != nil {
		return
	}

	// Each panel.
	// <package>/<panelName>Panel.go

	paneldata := &panelTemplateData{
		PackageName: packageName,
		Funcs:       utils.GetFuncs(),
	}
	for _, panel := range button.Panels {
		paneldata.Panel = panel
		fName = utils.DeCap(panel.Name) + "Panel.go"
		fPath = filepath.Join(panelGroupPackagePath, fName)
		if err = utils.ProcessTemplate(fName, fPath, panelTemplate, paneldata); err != nil {
			return
		}
	}

	return
}
