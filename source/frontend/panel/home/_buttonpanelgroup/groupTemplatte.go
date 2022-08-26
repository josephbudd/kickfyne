package buttonpanelgroup

import (
	"github.com/josephbudd/kickfyne/source/panelloader"
	"github.com/josephbudd/kickfyne/source/utils"
)

const (
	groupFileName = "group.go"
)

type groupTemplateData struct {
	GroupName        string
	PackageName      string
	ImportPrefix     string
	Button           panelloader.Button
	PanelNamesSorted []string
	Funcs            utils.Funcs
}

var groupTemplate = `{{ $DOT := . }}package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"{{ .ImportPrefix }}/frontend/panel"
)

var (
	groupID = panel.NextGroupID()
	groupName = "{{ .GroupName }}"
	groupContent *fyne.Container
	window fyne.Window
)
{{ range $panel := .Button.Panels }}

func show{{ $panel.Name }}Panel() {
	{{- range $panel2 := $DOT.Button.Panels }}
		{{- if ne $panel2 $panel}}
	{{ call $DOT.Funcs.DeCap $panel2.Name }}Panel.content.Hide()
		{{- end }}
	{{- end }}
	{{ call $DOT.Funcs.DeCap $panel.Name }}Panel.content.Show()
	groupContent.Refresh()
}
{{- end }}

// Init creates the content for each panel and starts the messenger.
func Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {
	messenger = &messageHandler{}
	stater = &stateHandler{}
	window = w

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.Init: %w", err)
		}
	}()

	// Build each panel.{{ range $panel := .Button.Panels }}
	build{{ $panel.Name }}Panel(){{ end }}

	groupContent = container.New(
		layout.NewMaxLayout(),{{ range $panel := .Button.Panels }}
		{{ call $DOT.Funcs.DeCap $panel.Name }}Panel.content,{{ end }}
	)
	show{{ (index .Button.Panels 0).Name }}Panel()

	// Start the stater so it's communicating with the state.
	stater.init()
	// Start the messenger so it's communicating with the back end.
	err = messenger.listen()
	return
}

// Content builds and returns the groups content.
func Content() (content *fyne.Container) {
	content = groupContent
	return
}

`
