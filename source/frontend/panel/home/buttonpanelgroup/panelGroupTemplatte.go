package buttonpanelgroup

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type groupTemplateData struct {
	HomeButtonName   string
	GroupName        string
	DefaultPanelName string
	PackageName      string
	ImportPrefix     string
	Funcs            utils.Funcs
}

const (
	groupFileName = "panelGroup.go"

	groupTemplate = `// Package {{ .PackageName }} is the {{ .HomeButtonName }} home button's panel group.
// State is read only and may effect how panels are updated.
// Messages received can effect how panels are updated.
// User input is relayed to the backend using messages sent.
package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"{{ .ImportPrefix }}/frontend/panel/builder"
)

type home{{ .HomeButtonName }}PanelGroup struct{}

const (
	homeButtonName = "{{ .HomeButtonName }}"
)

var (
	panelGroup *home{{ .HomeButtonName }}PanelGroup = &home{{ .HomeButtonName }}PanelGroup{}
)

func init() {
	builder.RegisterButtonPanelGroup(homeButtonName, panelGroup)
}

// Init .
// Init initializes this panel group.
// * It initialiizes each panel in this panel group.
// * It starts the stater.
// * It starts the messenger.
// Returns the error.
func (bldr *home{{ .HomeButtonName }}PanelGroup) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home{{ .HomeButtonName }}PanelGroup.Init: %w", err)
		}
	}()

	// Get the panel builders.
	panelBuilders := builder.ButtonPanels(homeButtonName)
	for _, panelBuilder := range panelBuilders {
		panelBuilder.Init(ctx, ctxCancel, app, w)
	}

	// Start the stater so it's communicating with the state.
	stater.listen()
	// Start the messenger so it's communicating with the back end.
	err = messenger.listen()
	return
}

// PanelGroup creates this panel group's content and then returns it.
func (bldr *home{{ .HomeButtonName }}PanelGroup) PanelGroup() (content fyne.CanvasObject) {
	// Get the panel builders.
	panelBuilders := builder.ButtonPanels(homeButtonName)
	contents := make([]fyne.CanvasObject, len(panelBuilders))
	for i, panelBuilder := range panelBuilders {
		contents[i] = panelBuilder.Panel()
	}
	content = container.NewVBox(contents...)
	{{ call .Funcs.DeCap .DefaultPanelName }}Panel.Show()
	return
}

`
)
