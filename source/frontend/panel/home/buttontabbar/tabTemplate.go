package buttontabbar

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type tabTemplateData struct {
	TabName                string
	TabLabel               string
	TabIndex               int
	PackageName            string
	HomeButtonName         string
	PanelGroupFolderImport string
	ImportPrefix           string
	Funcs                  utils.Funcs
}

func tabFileName(tabName string) (fileName string) {
	fileName = utils.DeCap(tabName) + "Tab.go"
	return
}

const tabTemplate = `package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"{{ .ImportPrefix }}/frontend/panel/builder"

	_ "{{ .PanelGroupFolderImport }}"
)

const (
	{{ call .Funcs.DeCap .TabName }}TabName  = "{{ .TabName }}"
	{{ call .Funcs.DeCap .TabName }}TabLabel = "{{ .TabLabel }}"
)

var (
	{{ call .Funcs.DeCap .TabName }}TabItemBuilder = &{{ .TabName }}TabItemBuilder{}
)

type {{ .TabName }}TabItemBuilder struct{}

func init() {
	builder.RegisterTabItem(
		homeButtonName,
		{{ .TabIndex }},
		{{ call .Funcs.DeCap .TabName }}TabItemBuilder,
	)
}

// Implement the builder.TabItemBuilder interface with funcs Init, TabItem and Name.
	
func (b *{{ .TabName }}TabItemBuilder) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .TabName }}TabItemBuilder.Init: %w", err)
		}
	}()

	panelGroupBuilder := builder.TabPanelGroup(homeButtonName, {{ call .Funcs.DeCap .TabName }}TabName)
	panelGroupBuilder.Init(ctx, ctxCancel, app, w)
	return
}

func (b *{{ .TabName }}TabItemBuilder) TabItem() (tabItem *container.TabItem) {
	panelGroupBuilder := builder.TabPanelGroup(homeButtonName, {{ call .Funcs.DeCap .TabName }}TabName)
	tabItem = container.NewTabItem({{ call .Funcs.DeCap .TabName }}TabLabel, panelGroupBuilder.PanelGroup())
	return
}

func (b *{{ .TabName }}TabItemBuilder) Name() (name string) {
	name = {{ call .Funcs.DeCap .TabName }}TabName
	return
}

`
