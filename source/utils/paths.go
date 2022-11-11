package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	folderNameBackend       = "backend"
	folderNameFrontend      = "frontend"
	folderNameShared        = "shared"
	FolderNameSpawn         = "spawn"
	folderNameGUI           = "gui"
	FolderNameLandingScreen = "landingscreen"
	folderNameMainMenu      = "mainmenu"
	FolderNameScreens       = "screens"
	folderNameMessage       = "message"
	folderNameStore         = "store"
	folderNameStoring       = "storing"
	folderNameStorer        = "storer"
	folderNameRecord        = "record"
	folderNameTXRX          = "txrx"
)

var (
	backendPaths        = filepath.Join(folderNameBackend, "folder")
	backendStore        = filepath.Join(folderNameBackend, folderNameStore)
	backendStoreStorer  = filepath.Join(backendStore, folderNameStorer)
	backendStoreStoring = filepath.Join(backendStore, folderNameStoring)
	backendTXRX         = filepath.Join(folderNameBackend, folderNameTXRX)

	frontendGUI         = filepath.Join(folderNameFrontend, folderNameGUI)
	frontendGUIMainMenu = filepath.Join(frontendGUI, folderNameMainMenu)
	frontendGUIScreens  = filepath.Join(frontendGUI, FolderNameScreens)
	frontendLanding     = filepath.Join(folderNameFrontend, FolderNameLandingScreen)
	frontendTXRX        = filepath.Join(folderNameFrontend, folderNameTXRX)

	sharedMessage     = filepath.Join(folderNameShared, folderNameMessage)
	sharedMetaData    = filepath.Join(folderNameShared, "meta")
	sharedStoreRecord = filepath.Join(folderNameShared, folderNameRecord)
)

type FolderPaths struct {
	App                                                   string
	Backend                                               string
	BackendPaths                                          string
	BackendStore, BackendStoreStorer, BackendStoreStoring string
	BackendTXRX                                           string

	Frontend                                             string
	FrontendGUI, FrontendGUIMainMenu, FrontendGUIScreens string
	FrontendLanding                                      string
	FrontendTXRX                                         string

	Shared                         string
	SharedMessage                  string
	SharedMetaData                 string
	SharedStore, SharedStoreRecord string
}

// BuildFolderPaths constructs paths and then makes them on the disk.
func BuildFolderPaths(rootPath string) (folderPaths *FolderPaths, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.BuildFolderPaths: %w", err)
		}
	}()

	folderPaths = &FolderPaths{
		App: rootPath,

		Backend:             filepath.Join(rootPath, folderNameBackend),
		BackendPaths:        filepath.Join(rootPath, backendPaths),
		BackendStore:        filepath.Join(rootPath, backendStore),
		BackendStoreStorer:  filepath.Join(rootPath, backendStoreStorer),
		BackendStoreStoring: filepath.Join(rootPath, backendStoreStoring),
		BackendTXRX:         filepath.Join(rootPath, backendTXRX),

		Frontend:            filepath.Join(rootPath, folderNameFrontend),
		FrontendGUI:         filepath.Join(rootPath, frontendGUI),
		FrontendGUIMainMenu: filepath.Join(rootPath, frontendGUIMainMenu),
		FrontendGUIScreens:  filepath.Join(rootPath, frontendGUIScreens),
		FrontendLanding:     filepath.Join(rootPath, frontendLanding),
		FrontendTXRX:        filepath.Join(rootPath, frontendTXRX),

		Shared:            filepath.Join(rootPath, folderNameShared),
		SharedMessage:     filepath.Join(rootPath, sharedMessage),
		SharedMetaData:    filepath.Join(rootPath, sharedMetaData),
		SharedStoreRecord: filepath.Join(rootPath, sharedStoreRecord),
	}
	err = buildFolderPaths(folderPaths)
	return
}

// RebuildFolderPaths remakes the folder paths on disk.
// Useful for restarting the framework.
func RebuildFolderPaths(folderPaths *FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.RebuildFolderPaths: %w", err)
		}
	}()

	err = buildFolderPaths(folderPaths)
	return
}

// buildFolderPaths constructs the paths onto the disk.
func buildFolderPaths(folderPaths *FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.buildFolderPaths: %w", err)
		}
	}()

	var isBuilt bool
	if isBuilt, err = IsBuilt(folderPaths.App); err != nil || isBuilt {
		// The folders have already been created.
		return
	}

	// Create the folders.

	// Backend.
	if err = os.Mkdir(folderPaths.Backend, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendPaths, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendStore, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendStoreStorer, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendStoreStoring, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendTXRX, DMode); err != nil {
		return
	}

	// Frontend.
	if err = os.Mkdir(folderPaths.Frontend, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendGUI, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendGUIMainMenu, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendGUIScreens, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendLanding, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendTXRX, DMode); err != nil {
		return
	}

	// Shared
	if err = os.Mkdir(folderPaths.Shared, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedMessage, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedMetaData, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedStoreRecord, DMode); err != nil {
		return
	}
	return

}

// UnBuildFolderPaths removes backend, frontend and shared folders.
func UnBuildFolderPaths(folderPaths *FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.UnBuildFolderPaths: %w", err)
		}
	}()

	// Remove the folders.

	// Backend.
	if err = os.RemoveAll(folderPaths.Backend); err != nil {
		return
	}

	// Frontend.
	if err = os.RemoveAll(folderPaths.Frontend); err != nil {
		return
	}

	// Shared
	err = os.RemoveAll(folderPaths.Shared)
	return

}
