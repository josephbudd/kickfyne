package main

import (
	"context"
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

	go notify(ctx, ctxCancel)
	dumperCh := startDumper(ctx)
	defer func() {
		dumperCh <- "EOJ\n"
	}()

	// doit(dumperCh)

	var pathWD string
	if pathWD, err = os.Getwd(); err != nil {
		return
	}

	var isBuilt bool
	var importPrefix string
	var folderPaths *utils.FolderPaths
	if isBuilt, importPrefix, folderPaths, err = utils.Setup(pathWD, dumperCh); err != nil {
		return
	}
	// Build the args to pass on to the handlers.
	var handlerArgs []string
	lArgs := len(os.Args)
	switch {
	case lArgs < 2:
		dumperCh <- help.Usage()
		return
	case lArgs > 2:
		handlerArgs = os.Args[2:]
	}

	switch os.Args[1] {
	case framework.Cmd:
		err = framework.Handler(pathWD, dumperCh, handlerArgs, isBuilt, importPrefix, folderPaths)
	case frontend.Cmd:
		err = frontend.Handler(pathWD, dumperCh, handlerArgs, isBuilt, importPrefix, folderPaths)
	case help.Cmd:
		err = help.Handler(pathWD, dumperCh, handlerArgs, isBuilt, importPrefix, folderPaths)
	case message.Cmd:
		err = message.Handler(pathWD, dumperCh, handlerArgs, isBuilt, importPrefix, folderPaths)
	case record.Cmd:
		err = record.Handler(pathWD, dumperCh, handlerArgs, isBuilt, importPrefix, folderPaths)
	default:
		dumperCh <- help.Usage()
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
