package source

import (
	"fmt"

	"github.com/josephbudd/kickfyne/source/shared/message"
	"github.com/josephbudd/kickfyne/source/utils"
)

// RemoveMessage removes a message.
func RemoveMessage(
	messageName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.RemoveMessage: %w", err)
		}
	}()

	err = message.RemoveMessage(messageName, folderPaths)

	return
}

// AddMessage adds a message.
func AddMessage(
	messageName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.AddMessage: %w", err)
		}
	}()

	err = message.AddMessage(messageName, folderPaths)

	return
}
