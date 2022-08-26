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

	"{{ .ImportPrefix }}/frontend/panel/home"
	"{{ .ImportPrefix }}/frontend/txrx"
	"{{ .ImportPrefix }}/shared/message"
)

func Start(ctx context.Context, ctxCancel context.CancelFunc, a fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("frontend.Start: %w", err)
		}
	}()

	// Initialize the panel groups.
	if err = home.Init(ctx, ctxCancel, a, w); err != nil {
		return
	}

	// Start communications with the back end.
	// The listener will run as a concurrent process.
	txrx.Listen(ctx, ctxCancel)

	// Send the init message.
	// Let the back end know that the front end is ready.
	// Parts of the front need data from the back end to build panel content.
	message.FrontEndToBackEnd <- message.NewInit()
	return
}

// Content builds and returns the view content.
// One panelgroup at a time.
func Content() (content *fyne.Container, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("renderer.Content: %w", err)
		}
	}()

	content, err = home.Content()
	return
}

`
