package message

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source"
	"github.com/josephbudd/kickfyne/source/backend/txrx"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd = "message"

	verbAdd    = "add"
	verbRemove = "remove"
	verbList   = "list"
	verbHelp   = "help"
)

func Handler(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.Handler: %w", err)
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
			dumperCh <- "The app must be initailized before a message can be added."
			return
		}
		if len(args) < 2 {
			dumperCh <- "Missing the message name."
			dumperCh <- Usage()
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = utils.ValidateNewMessageName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			dumperCh <- errMessage
			return
		}
		// Add a shared message.
		if err = source.AddMessage(args[1], folderPaths); err != nil {
			return
		}
		// Add a back end message handler.
		if err = txrx.AddMessageHandler(args[1], importPrefix, folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("Success. Message named %q added.", args[1])
	case verbRemove:
		if !isBuilt {
			dumperCh <- "The app must be initailized before a message can be removed."
			return
		}
		if len(args) < 2 {
			dumperCh <- "Missing the message name."
			dumperCh <- Usage()
			return
		}
		var isValid bool
		var errMessage string
		if isValid, errMessage, err = utils.ValidateCurrentMessageName(args[1], folderPaths); err != nil {
			return
		}
		if !isValid {
			dumperCh <- errMessage
			return
		}
		// Remove a message.
		if err = source.RemoveMessage(args[1], folderPaths); err != nil {
			return
		}
		// Remove a back end message handler.
		if err = txrx.RemoveMessageHandler(args[1], folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("Success. Message named %q removed.", args[1])
	case verbList:
		if !isBuilt {
			dumperCh <- "The app must be initailized before a message names can be listed."
			return
		}
		// List all of the messages.
		var messageNames []string
		if messageNames, err = utils.UserMessageNames(folderPaths); err != nil {
			return
		}
		dumperCh <- fmt.Sprintf("There are %d message names:\n", len(messageNames))
		for i, messageName := range messageNames {
			j := i + 1
			switch {
			case j < 10:
				dumperCh <- fmt.Sprintf("  %d  %s\n", j, messageName)
			default:
				dumperCh <- fmt.Sprintf("  %d %s\n", j, messageName)
			}
		}
	case verbHelp:
		dumperCh <- Usage()
	default:
		dumperCh <- fmt.Sprintf("\nUnknown command %q.\n", args[0])
		dumperCh <- Usage()
	}
	return
}
