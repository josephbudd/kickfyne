package tabpanelgroup

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	groupFileName = "panelGroup.go"
)

type groupTemplateData struct {
	HomeButtonName   string
	GroupName        string
	TabName          string
	DefaultPanelName string
	PackageName      string
	ImportPrefix     string
	Funcs            utils.Funcs
}

var groupTemplate = `// Package {{ .PackageName }} is the {{ .TabName }} tab's panel group.
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

type home{{ .HomeButtonName }}{{ .TabName }}Group struct{}

const (
	homeButtonName = "{{ .HomeButtonName }}"
	tabName        = "{{ .TabName }}"
)

var (
	panelGroup *home{{ .HomeButtonName }}{{ .TabName }}Group = &home{{ .HomeButtonName }}{{ .TabName }}Group{}
)

func init() {
	builder.RegisterTabPanelGroup(homeButtonName, tabName, panelGroup)
}

// Init .
// Init initializes this panel group.
// * It initialiizes each panel in this panel group.
// * It starts the messenger.
// Returns the error.
func (bldr *home{{ .HomeButtonName }}{{ .TabName }}Group) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("home{{ .HomeButtonName }}{{ .TabName }}Group.Init: %w", err)
		}
	}()

	// Get the panel builders.
	panelBuilders := builder.TabPanels(homeButtonName, tabName)
	for _, panelBuilder := range panelBuilders {
		panelBuilder.Init(ctx, ctxCancel, app, w)
	}

	// Start the messenger so it's communicating with the back end.
	err = messenger.listen()
	return
}

// PanelGroup creates this panel group's content and then returns it.
func (bldr *home{{ .HomeButtonName }}{{ .TabName }}Group) PanelGroup() (content fyne.CanvasObject) {
	// Get the panel builders.
	panelBuilders := builder.TabPanels(homeButtonName, tabName)
	contents := make([]fyne.CanvasObject, len(panelBuilders))
	for i, panelBuilder := range panelBuilders {
		contents[i] = panelBuilder.Panel()
	}
	content = container.NewVBox(contents...)
	{{ call .Funcs.DeCap .DefaultPanelName }}Panel.Show()
	return
}

`
