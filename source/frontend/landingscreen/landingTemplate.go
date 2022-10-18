package landingscreen

type landingTemplateData struct {
	ImportPrefix      string
	LandingScreenName string
}

const (
	landingFileName = "landing.go"

	landingTemplate = `package landingscreen


import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	"{{ .ImportPrefix }}/frontend/gui"
	"{{ .ImportPrefix }}/frontend/gui/screens/{{ .LandingScreenName }}"
)

var landingScreen gui.CanvasObjectProvider
var landingScreenErr error

// Init binds the landing screen to the app window.
func Init(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if landingScreenErr != nil {
			landingScreenErr = fmt.Errorf("landingscreen.SetLanding: %w", landingScreenErr)
		}
		err = landingScreenErr
	}()

	// Show the landing screen.
	if landingScreen, landingScreenErr = {{ .LandingScreenName }}.New(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}
	landingScreen.BindToWindow()
	return
}

// Land binds the landing to the window.
func Land() {
	if landingScreenErr != nil {
		return
	}
	landingScreen.BindToWindow()
}

`
)
