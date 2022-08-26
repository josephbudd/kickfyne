package message

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/josephbudd/kickfyne/source/utils"
)

func AddMessage(
	messageName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.Add: %w", err)
		}
	}()

	fname := utils.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.SharedMessage, fname)
	data := messageTemplateData{
		MessageName: messageName,
	}
	err = utils.ProcessTemplate(fname, oPath, messageTemplate, data)
	return
}

func RemoveMessage(
	messageName string,
	folderPaths *utils.FolderPaths,
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("message.RemoveMessage: %w", err)
		}
	}()

	fName := utils.MessageFileName(messageName)
	oPath := filepath.Join(folderPaths.SharedMessage, fName)
	err = os.Remove(oPath)
	return
}
