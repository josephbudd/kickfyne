package screens

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type screenTemplateData struct {
	PackageName      string
	PanelNames       []string
	DefaultPanelName string
	ImportPrefix     string
	Funcs            utils.Funcs
}

const (
	screenTemplate = `{{ $DOT := . -}}
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
{{- if eq (len .PanelNames) 0 }}
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
{{- end }}

	"{{ .ImportPrefix }}/frontend/gui"
)

// screenComponents is this screen, it's panels and messenger.
// This screen has {{ len .PanelNames }} panels.
{{- if ne (len .DefaultPanelName ) 0 }}
// The default panel is {{ .DefaultPanelName }}.
{{- end }}
type screenComponents struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	app       fyne.App
	window    fyne.Window

	canvasObjectProvider gui.CanvasObjectProvider
{{- if ne (len .PanelNames) 0 }}
	panels               *panels
{{- end}}
	messenger            *messageHandler
}
{{- if ne (len .PanelNames) 0 }}

type panels struct {
 {{- range $panelName := .PanelNames }}
	{{ $panelName }}   *{{ $panelName }}Components
 {{- end}}
}
{{- end}}

var packageScreen *screenComponents
var initErr error

// New returns the screen's canvas object provider and the error. 
// If needed, it constructs this screen.
// * It constructs each panel in this screen.
// * It constructs the messenger.
// * It uses the default panel contents to create the screen content.
// If this screen has already been constructed then it uses the already constructed screen.
func New(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (canvasObjectProvider gui.CanvasObjectProvider, err error) {

	if packageScreen != nil {
		canvasObjectProvider = packageScreen.canvasObjectProvider
		err = initErr
		return
	}

	var newScreen *screenComponents
	defer func() {
		if err == nil {
			packageScreen = newScreen
			canvasObjectProvider = packageScreen.canvasObjectProvider
		}
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.New: %w", err)
			initErr = err
		}
	}()

	newScreen = &screenComponents{
		ctx:                  ctx,
		ctxCancel:            ctxCancel,
		app:                  app,
		window:               w,
		canvasObjectProvider: gui.NewScreenCanvasManager(),
{{- if ne (len .PanelNames) 0 }}
		panels:               &panels{},
{{- end}}
		}
	canvasObjectProvider = newScreen.canvasObjectProvider

	// Make sure the canvasObjectProvider is constructed.
{{- if ne (len .PanelNames) 0 }}

	// Construct each panel building panel content.
	// Do the default panel last.
 {{- range $panelName := .PanelNames }}
  {{- if ne $panelName $DOT.DefaultPanelName }}
	var {{ $panelName }} *{{ $panelName }}Components
	if {{ $panelName }}, err = new{{ call $DOT.Funcs.Cap $panelName }}Components(newScreen); err != nil {
		return
	}
	newScreen.panels.{{ $panelName }} = {{ $panelName }}
  {{- end }}
 {{- end }}
 {{- if ne (len .DefaultPanelName) 0 }}
	var {{ .DefaultPanelName }} *{{ .DefaultPanelName }}Components
	if {{ .DefaultPanelName }}, err = new{{ call .Funcs.Cap .DefaultPanelName }}Components(newScreen); err != nil {
		return
	}
	newScreen.panels.{{ .DefaultPanelName }} = {{ .DefaultPanelName }}
	newScreen.canvasObjectProvider.UpdateCanvasObject({{ .DefaultPanelName }}.content)
 {{- end }}
{{- else }}

	// There are no panels in this screen package so display the screen package name.
	content := container.NewVBox(
		widget.NewLabel("This is the {{ .PackageName }} screen package."),
		widget.NewLabel("This screen package does not have any panels."),
	)
	newScreen.canvasObjectProvider.UpdateCanvasObject(content)
{{- end }}
	// Messenger.
	newScreen.messenger = newMessageHandler(newScreen)
	return
}

`
)
