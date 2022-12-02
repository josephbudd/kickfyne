package root

import "github.com/josephbudd/kickfyne/source/utils"

const (
	MainFileName = "main.go"
)

type mainTemplateData struct {
	ImportPrefix       string
	AppName            string
	ScreenPackageNames []string
	HomePackageName    string
	Funcs              utils.Funcs
}

var mainTemplate = `{{ $DOT := . }}package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"{{ .ImportPrefix }}/backend/store"
	betxrx "{{ .ImportPrefix }}/backend/txrx"
	"{{ .ImportPrefix }}/frontend/gui"
	"{{ .ImportPrefix }}/frontend/gui/mainmenu"
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

	// Size and position the window.
	if !fyne.CurrentDevice().IsMobile() {
		w.Resize(size16x9(1000, 0))
		w.CenterOnScreen()
	} else {
		w.SetFullScreen(true)
	}
	w.Show()

	// Start the back-end.
	// backend.Start also opens and returns the stores.
	// func waitAndClose will close the stores when the applicatioinlication ends.
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

	defer func() {
		if err != nil {
			err = fmt.Errorf("main.backendStart: %w", err)
		}
	}()

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
func frontendStart(ctx context.Context, ctxCancelFunc context.CancelFunc, application fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("main.frontendStart: %w", err)
		}
	}()

	// Initialize the view.
	gui.Init(window)

	// Build the screens.
	screenStart(ctx, ctxCancelFunc, application, window)
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
		clErr = stores.CloseTimeout(1)
		w.Close()
		return
	case err = <-errCh:
		clErr = stores.CloseTimeout(1)
		w.Close()
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

// screenStart starts the screens.
func screenStart(ctx context.Context, ctxCancelFunc context.CancelFunc, application fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("main.frontendStart: %w", err)
		}
	}()

	// Initialize main menu.
	if err = mainmenu.Init(ctx, ctxCancelFunc, application, window); err != nil {
		return
	}

	screenLoadedCh := make(chan gui.CanvasObjectProvider, countScreens)
	if err = startScreens(ctx, ctxCancelFunc, application, window, screenLoadedCh); err != nil {
		return
	}

	// Start communications with the back-end.
	// The receiver will run as a concurrent process.
	fetxrx.StartReceiver(ctx, ctxCancelFunc)

	// Wait for the screens to load and then add the main menu.
	go waitToAddMainMenu(ctx, screenLoadedCh)
	return
}

func waitToAddMainMenu(ctx context.Context, screenLoadedCh chan gui.CanvasObjectProvider) {
	// Wait for the all the screens to load.
	var countLoaded int
	nScreens := cap(screenLoadedCh)
	for countLoaded < nScreens {
		select {
		case <-screenLoadedCh:
			countLoaded++
		case <-ctx.Done():
			return
		}
	}
	// All the screens are loaded.
	// Display the main menu.
	mainmenu.Reset(nil)
}

`
