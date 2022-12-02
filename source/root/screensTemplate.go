package root

import "github.com/josephbudd/kickfyne/source/utils"

const (
	ScreensFileName = "screens.go"
)

type screensTemplateData struct {
	ImportPrefix       string
	AppName            string
	ScreenPackageNames []string
	HomePackageName    string
	Funcs              utils.Funcs
}

var screensTemplate = `{{ $DOT := . }}package main

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	"{{ .ImportPrefix }}/frontend/gui"
{{- range $screenName := .ScreenPackageNames }}
	"{{ $DOT.ImportPrefix }}/frontend/gui/screens/{{ $screenName }}"
{{- end }}
	"{{ .ImportPrefix }}/frontend/landingscreen"
)

const (
	countScreens = {{ len .ScreenPackageNames }}
)

// startScreens starts the screens.
func startScreens(ctx context.Context, ctxCancelFunc context.CancelFunc, application fyne.App, window fyne.Window, screenLoadedCh chan gui.CanvasObjectProvider) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("main.startScreens: %w", err)
		}
	}()

	// Build each screen's gui.CanvasObjectProvider.
{{- range $screenName := .ScreenPackageNames }}
	// {{ $screenName }}.
	if err = {{ $screenName }}.New(ctx, ctxCancelFunc, application, window, screenLoadedCh); err != nil {
		return
	}
{{- end }}

	// Display the landing screen.
	landingscreen.Init()
	return
}

`
