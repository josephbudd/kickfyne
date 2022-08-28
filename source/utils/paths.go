package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	FolderNameBackend  = "backend"
	FolderNameFrontend = "frontend"
	FolderNameShared   = "shared"
)

var (
	backendMisc        = filepath.Join(FolderNameBackend, "misc")
	backendMiscShuffle = filepath.Join(backendMisc, "shuffle")
	backendTXRX        = filepath.Join(FolderNameBackend, "txrx")

	frontendPanel            = filepath.Join(FolderNameFrontend, "panel")
	frontendPanelBuilder     = filepath.Join(frontendPanel, "builder")
	frontendPanelHome        = filepath.Join(frontendPanel, "home")
	frontendWidget           = filepath.Join(FolderNameFrontend, "widget")
	frontendWidgetBackPanel  = filepath.Join(frontendWidget, "backpanel")
	frontendWidgetSafeButton = filepath.Join(frontendWidget, "safebutton")
	frontendWidgetSelection  = filepath.Join(frontendWidget, "selection")
	frontendTXRX             = filepath.Join(FolderNameFrontend, "txrx")

	sharedMessage      = filepath.Join(FolderNameShared, "message")
	sharedPaths        = filepath.Join(FolderNameShared, "paths")
	sharedStore        = filepath.Join(FolderNameShared, "store")
	sharedStoreRecord  = filepath.Join(sharedStore, "record")
	sharedStoreStorer  = filepath.Join(sharedStore, "storer")
	sharedStoreStoring = filepath.Join(sharedStore, "storing")
)

type FolderPaths struct {
	App                             string
	Backend                         string
	BackendMisc, BackendMiscShuffle string
	BackendTXRX                     string

	Frontend                                                                                   string
	FrontendPanel, FrontendPanelBuilder, FrontendPanelHome                                     string
	FrontendWidget, FrontendWidgetBackPanel, FrontendWidgetSafeButton, FrontendWidgetSelection string
	FrontendTXRX                                                                               string

	Shared                                                                string
	SharedMessage                                                         string
	SharedPaths                                                           string
	SharedStore, SharedStoreRecord, SharedStoreStorer, SharedStoreStoring string
}

func BuildFolderPaths(rootPath string) (folderPaths *FolderPaths, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.BuildFolderPaths: %w", err)
		}
	}()

	folderPaths = &FolderPaths{
		App: rootPath,

		Backend:            filepath.Join(rootPath, FolderNameBackend),
		BackendMisc:        filepath.Join(rootPath, backendMisc),
		BackendMiscShuffle: filepath.Join(rootPath, backendMiscShuffle),
		BackendTXRX:        filepath.Join(rootPath, backendTXRX),

		Frontend:                 filepath.Join(rootPath, FolderNameFrontend),
		FrontendPanel:            filepath.Join(rootPath, frontendPanel),
		FrontendPanelBuilder:     filepath.Join(rootPath, frontendPanelBuilder),
		FrontendPanelHome:        filepath.Join(rootPath, frontendPanelHome),
		FrontendWidget:           filepath.Join(rootPath, frontendWidget),
		FrontendWidgetBackPanel:  filepath.Join(rootPath, frontendWidgetBackPanel),
		FrontendWidgetSafeButton: filepath.Join(rootPath, frontendWidgetSafeButton),
		FrontendWidgetSelection:  filepath.Join(rootPath, frontendWidgetSelection),
		FrontendTXRX:             filepath.Join(rootPath, frontendTXRX),

		Shared:             filepath.Join(rootPath, FolderNameShared),
		SharedMessage:      filepath.Join(rootPath, sharedMessage),
		SharedPaths:        filepath.Join(rootPath, sharedPaths),
		SharedStore:        filepath.Join(rootPath, sharedStore),
		SharedStoreRecord:  filepath.Join(rootPath, sharedStoreRecord),
		SharedStoreStorer:  filepath.Join(rootPath, sharedStoreStorer),
		SharedStoreStoring: filepath.Join(rootPath, sharedStoreStoring),
	}

	var isBuilt bool
	if isBuilt, err = gotBuilt(rootPath); err != nil || isBuilt {
		// The folders have already been created.
		return
	}

	// Create the folders.

	// Backend.
	if err = os.Mkdir(folderPaths.Backend, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendMisc, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendMiscShuffle, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.BackendTXRX, DMode); err != nil {
		return
	}

	// Frontend.
	if err = os.Mkdir(folderPaths.Frontend, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendPanel, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendPanelBuilder, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendPanelHome, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendWidget, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendWidgetBackPanel, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendWidgetSafeButton, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.FrontendWidgetSelection, DMode); err != nil {
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
	if err = os.Mkdir(folderPaths.SharedPaths, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedStore, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedStoreRecord, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedStoreStorer, DMode); err != nil {
		return
	}
	if err = os.Mkdir(folderPaths.SharedStoreStoring, DMode); err != nil {
		return
	}
	return
}

func (fp *FolderPaths) ModDotGoPath() (path string) {
	path = filepath.Join(fp.App, "mod.go")
	return
}
