package help

import (
	"fmt"

	"github.com/josephbudd/kickfyne/commands/frontend"
	"github.com/josephbudd/kickfyne/commands/message"
	"github.com/josephbudd/kickfyne/commands/record"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	Cmd         = "help"
	usageSubCmd = "usage"
)

func Handler(pathWD string, dumperCh chan string, args []string, isBuilt bool, importPrefix string, folderPaths *utils.FolderPaths) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("help.Handler: %w", err)
		}
	}()

	if len(args) == 0 {
		dumperCh <- Usage()
		return
	}

	switch args[0] {
	case frontend.Cmd:
		dumperCh <- frontend.Usage()
	case message.Cmd:
		dumperCh <- message.Usage()
	case record.Cmd:
		dumperCh <- record.Usage()
	case usageSubCmd:
		dumperCh <- Usage()
	default:
		Usage()
	}
	return
}
