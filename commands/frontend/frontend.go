package frontend

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd          = "frontend"
	SubCmdCreate = "create"
	subCmdClear  = "clear"
	SubCmdButton = "button"
	SubCmdTab    = "tab"
	SubCmdPanel  = "panel"
	subCmdHelp   = "help"

	verbAdd    = "add"
	verbRemove = "remove"
	verbList   = "list"
)

func Handler(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Handler: %w", err)
		}
	}()

	if len(args) == 0 {
		dumperCh <- "Missing a sub command."
		dumperCh <- Usage()
		return
	}
	switch args[0] {
	case SubCmdCreate:
		err = handleCreate(pathWD, dumperCh, args, isBuilt, importPrefix, folderPaths)
	case subCmdClear:
		err = handleClear(dumperCh, isBuilt, folderPaths)
	case SubCmdButton:
		err = handleButton(pathWD, dumperCh, args, isBuilt, importPrefix, folderPaths)
	case SubCmdTab:
		err = handleTab(pathWD, dumperCh, args, isBuilt, importPrefix, folderPaths)
	case SubCmdPanel:
		err = handlePanel(pathWD, dumperCh, args, isBuilt, importPrefix, folderPaths)
	default:
		dumperCh <- fmt.Sprintf("\nUnknown command %q.\n", args[0])
		dumperCh <- Usage()
	}
	return
}
