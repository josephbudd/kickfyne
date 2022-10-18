package frontend

const (
	frontendFileName = "frontend.go"
)

type frontendTemplateData struct {
	ImportPrefix string
}

var frontendTemplate = `package frontend

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	"{{ .ImportPrefix }}/frontend/gui"
	"{{ .ImportPrefix }}/frontend/gui/mainmenu"
	"{{ .ImportPrefix }}/frontend/landingscreen"
	"{{ .ImportPrefix }}/frontend/txrx"
	"{{ .ImportPrefix }}/shared/message"
)

func Start(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Initialize the view stack
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
	txrx.StartReceiver(ctx, ctxCancelFunc)

	// Send the init message.
	// Let the back-end know that the front end is ready.
	// Parts of the front need data from the back-end to build panel content.
	message.FrontEndToBackEnd <- message.NewInit()
	return
}

`
