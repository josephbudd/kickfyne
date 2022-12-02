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
type screenComponents struct {
	ctx                  context.Context
	ctxCancel            context.CancelFunc
	application          fyne.App
	window               fyne.Window
	loadedCh             chan gui.CanvasObjectProvider
	loaded               bool
	canvasObjectProvider gui.CanvasObjectProvider
	panels               *panels
	messenger            *messageHandler
}

type panels struct {
{{- range $panelName := .PanelNames }}
	{{ $panelName }}   *{{ $panelName }}Components
{{- end}}
}

func (screen *screenComponents) signalLoaded() {
	if screen.loaded {
		return
	}
	screen.loaded = true
	screen.loadedCh <- screen.canvasObjectProvider
}

var packageScreen *screenComponents

// CanvasObjectProvider returns the screen's gui.CanvasObjectProvider.
func CanvasObjectProvider() (canvasObjectProvider gui.CanvasObjectProvider) {
	canvasObjectProvider = packageScreen.canvasObjectProvider
	return
}

// New constructs this screen.
// Returns the error.
func New(ctx context.Context, ctxCancel context.CancelFunc, application fyne.App, w fyne.Window, loadedCh chan gui.CanvasObjectProvider) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("courses.New: %w", err)
			return
		}

		/* Signal that the screen has loaded.

		   KICKFYNE TODO:
		   If the displaying of the initial content in panels may be delayed
		    you should not call packageScreen.signalLoaded() here.
		   Instead you could call it inside the message receive func loading the panel.
		   That's because the message dispather dispatches messages using go threads.
		   Call it after the last panel has loaded.
		*/
		packageScreen.signalLoaded()
	}()

	packageScreen = &screenComponents{
		ctx:                  ctx,
		ctxCancel:            ctxCancel,
		application:          application,
		window:               w,
		canvasObjectProvider: gui.NewScreenCanvasManager(),
		panels:               &panels{},
		loadedCh:             loadedCh,
	}

	// Make sure the canvasObjectProvider is constructed.
{{- if ne (len .PanelNames) 0 }}

	// Construct each panel building panel content.
 {{- range $panelName := .PanelNames }}
  {{- if ne $panelName $DOT.DefaultPanelName }}
    // {{ $DOT.Func.Cap $panelName }}
	var {{ $panelName }} *{{ $panelName }}Components
	if {{ $panelName }}, err = new{{ $DOT.Func.Cap $panelName }}Components(packageScreen); err != nil {
		return
	}
	packageScreen.panels.{{ $panelName }}Panel = {{ $panelName }}Panel
  {{- end }}
 {{- end }}
{{- else }}

	// There are no panels in this screen package so display the screen package name.
	content := container.NewVBox(
		widget.NewLabel("This is the {{ .PackageName }} screen package."),
		widget.NewLabel("This screen package does not have any panels."),
	)
	packageScreen.canvasObjectProvider.UpdateCanvasObject(content)
{{- end }}
	// Messenger.
	packageScreen.messenger = newMessageHandler(packageScreen)
	return
}

`
)
