package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/josephbudd/kickfyne/commands/framework"
	"github.com/josephbudd/kickfyne/commands/frontend"
	"github.com/josephbudd/kickfyne/commands/help"
	"github.com/josephbudd/kickfyne/commands/message"
	"github.com/josephbudd/kickfyne/commands/record"
	"github.com/josephbudd/kickfyne/source/utils"
)

func main() {

	var err error
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer func() {
		ctxCancel()
		if err != nil {
			log.Println("Error: ", err.Error())
			os.Exit(1)
		}
	}()

	var pathWD string
	if pathWD, err = os.Getwd(); err != nil {
		return
	}

	var isBuilt bool
	if isBuilt, err = utils.IsBuilt(pathWD); err != nil {
		return
	}
	var importPrefix string
	if importPrefix, err = utils.ImportPrefix(pathWD); err != nil {
		fmt.Println("Failure: Unable to read a go.mod file.")
		return
	}
	go notify(ctx, ctxCancel)
	// Build the args to pass on to the handlers.
	lArgs := len(os.Args)
	if lArgs < 2 {
		fmt.Println(help.Usage())
		return
	}

	switch os.Args[1] {
	case framework.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = framework.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case frontend.CmdScreen, frontend.CmdPanel:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[1:]
		}
		err = frontend.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case help.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = help.Handler(handlerArgs)
	case message.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = message.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case record.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = record.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	default:
		fmt.Println(help.Usage())
	}
}

func notify(ctx context.Context, ctxCancel context.CancelFunc) {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-signalChan:
			ctxCancel()
			// terminate after second signal before callback is done
			go func() {
				<-signalChan
				os.Exit(1)
			}()
			return
		}
	}
}
