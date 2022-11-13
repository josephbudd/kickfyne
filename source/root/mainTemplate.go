package root

import "github.com/josephbudd/kickfyne/source/utils"

const (
	MainFileName = "main.go"
)

type mainTemplateData struct {
	ImportPrefix string
	AppName      string
	Funcs        utils.Funcs
}

var mainTemplate = `{{ $lCAppName := call .Funcs.LowerCase .AppName }}package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	betxrx "{{ .ImportPrefix }}/backend/txrx"
	"{{ .ImportPrefix }}/backend/store"
	"{{ .ImportPrefix }}/frontend/gui"
	"{{ .ImportPrefix }}/frontend/gui/mainmenu"
	"{{ .ImportPrefix }}/frontend/landingscreen"
	fetxrx "{{ .ImportPrefix }}/frontend/txrx"
	"{{ .ImportPrefix }}/shared/message"
)

var (
	stores *store.Stores
)

func main() {

	var err error
	defer func() {
		if err != nil {
			log.Printf("main: err is %s", err.Error())
			os.Exit(1)
		}
	}()

	if len(os.Getenv("FYNE_SCALE")) == 0 {
		os.Setenv("FYNE_SCALE", "1")
	}
	if len(os.Getenv("FYNE_THEME")) == 0 {
		os.Setenv("FYNE_THEME", "dark")
	}

	a := app.New()
	w := a.NewWindow("{{ .AppName }}")

	// Cancel.
	ctx, ctxCancel := context.WithCancel(context.Background())
	w.SetCloseIntercept(
		ctxCancel,
	)
	errCh := make(chan error, 2)
	go waitAndClose(w, ctx, ctxCancel, errCh)

	// Start the front end.
	if err = frontendStart(ctx, ctxCancel, a, w); err != nil {
		return
	}

	size := size16x9(1000, 0)
	w.Resize(size)
	w.CenterOnScreen()
	w.Show()

	// Start the back-end.
	// backend.Start also opens and returns the stores.
	// func waitAndClose will close the stores when the app ends.
	if stores, err = backendStart(ctx, ctxCancel); err != nil {
		return
	}

	// Send the init message to the back-end letting it know that the front end is ready.
	// See backend/txrx/init.go for details on 
	// * completing any backend initializations.
	// * sending messages to the front-end with data for the panels to display.
	message.FrontEndToBackEnd <- message.NewInit()

	// Start Fyne's event cycle.
	a.Run()
}

// backendStart starts the backend.
func backendStart(ctx context.Context, ctxCancel context.CancelFunc) (stores *store.Stores, err error) {
	// Opens the stores.
	if stores, err = store.New(); err != nil {
		return
	}
	if err = stores.Open(); err != nil {
		return
	}
	// Receive messages from the front-end.
	betxrx.StartReceiver(ctx, ctxCancel, stores)
	return
}

// frontendStart starts the front-end.
func frontendStart(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Initialize the view.
	gui.Init(window)

	// Show the landing screen.
	if err = landingscreen.Init(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}

	// Initialize main menu.
	// The developer must ensure that all panel groups should get initialized from main menu.
	if err = mainmenu.Init(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}

	// Start communications with the back-end.
	// The receiver will run as a concurrent process.
	fetxrx.StartReceiver(ctx, ctxCancelFunc)
	return
}

// waitAndClose waits for the context to end and then closes down.
func waitAndClose(w fyne.Window, ctx context.Context, ctxCancel context.CancelFunc, errCh chan error) {

	var err error
	var clErr error
	defer func() {
		var level int
		if err != nil {
			level++
			log.Println(err)
		}
		if clErr != nil {
			level++
			log.Println(clErr)
		}
		os.Exit(level)
	}()

	select {
	case <-ctx.Done():
		w.Close()
		clErr = stores.Close()
		return
	case err = <-errCh:
		w.Close()
		clErr = stores.Close()
		return
	}
}

func size16x9(width, height int) (size fyne.Size) {
	var newWidth float32
	var newHeight float32
	switch {
	case width != 0:
		if width < 0 {
			width = 0 - width
		}
		r := width / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	case height != 0:
		if height < 0 {
			height = 0 - height
		}
		r := height / 9
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	default:
		// default to 720 width.
		r := 720 / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	}
	size = fyne.Size{Width: newWidth, Height: newHeight}
	return
}

`
