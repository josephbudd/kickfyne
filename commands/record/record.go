package record

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd = "record"

	verbAdd    = "add"
	verbRemove = "remove"
	verbList   = "list"
	verbHelp   = "help"
)

func Handler(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("record.Handler: %w", err)
		}
	}()

	if len(args) == 0 {
		dumperCh <- "Missing the verb."
		dumperCh <- Usage()
		return
	}
	switch args[0] {
	case verbAdd:
		if !isBuilt {
			dumperCh <- "The app must be initailized before a record can be added."
			return
		}
		if len(args) < 2 {
			dumperCh <- "Missing the record name."
			dumperCh <- Usage()
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = utils.ValidateNewRecordName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			dumperCh <- errMessage
			return
		}
		// Add a record.
		if err = source.AddRecord(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("Success. Record named %q added.", args[1])
	case verbRemove:
		if !isBuilt {
			dumperCh <- "The app must be initailized before a record can be removed."
			return
		}
		if len(args) < 2 {
			dumperCh <- "Missing the record name."
			dumperCh <- Usage()
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = utils.ValidateCurrentRecordName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			dumperCh <- errMessage
			return
		}
		// Remove a record.
		if err = source.RemoveRecord(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("Success. Record named %q removed.", args[1])
	case verbList:
		if !isBuilt {
			dumperCh <- "The app must be initailized before a record names can be listed."
			return
		}
		// List all of the records.
		var recordNames []string
		if recordNames, err = utils.UserRecordNames(folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("There are %d record names:\n", len(recordNames))
		for i, recordName := range recordNames {
			j := i + 1
			switch {
			case j < 10:
				dumperCh <- fmt.Sprintf("  %d  %s\n", j, recordName)
			default:
				dumperCh <- fmt.Sprintf("  %d %s\n", j, recordName)
			}
		}
	case verbHelp:
		dumperCh <- Usage()
	default:
		dumperCh <- fmt.Sprintf("\nUnknown command %q.\n", args[0])
	}
	return
}
