package help

import (
	"fmt"

	"github.com/josephbudd/kickfyne/commands/frontend"
	"github.com/josephbudd/kickfyne/commands/message"
	"github.com/josephbudd/kickfyne/commands/record"
)

const (
	Cmd = "help"
)

// Handler displays the requested help.
func Handler(args []string) (err error) {

	if len(args) == 0 {
		fmt.Println(Usage())
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("help.Handler: %w", err)
		}
	}()

	switch args[0] {
	case frontend.CmdScreen:
		fmt.Println(frontend.UsageScreen())
	case frontend.CmdPanel:
		fmt.Println(frontend.UsagePanel())
	case message.Cmd:
		fmt.Println(message.Usage())
	case record.Cmd:
		fmt.Println(record.Usage())
	default:
		Usage()
	}
	return
}
