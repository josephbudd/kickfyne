package home

import "github.com/josephbudd/kickfyne/source/utils"

type buttonTemplateData struct {
	ImportPrefix            string
	ButtonName              string
	ButtonLabel             string
	ButtonIndex             int
	ButtonGroupFolderImport string
	Funcs                   utils.Funcs
}

const (
	buttonTemplate = `package home

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"{{ .ImportPrefix }}/frontend/panel/builder"
	"{{ .ImportPrefix }}/frontend/widget/backpanel"
	"{{ .ImportPrefix }}/frontend/widget/safebutton"

	_ "{{ .ButtonGroupFolderImport }}"
)

const (
	{{ call .Funcs.DeCap .ButtonName }}ButtonName  = "{{ .ButtonName }}"
	{{ call .Funcs.DeCap .ButtonName }}ButtonLabel = "{{ .ButtonLabel }}"
)

var (
	{{ call .Funcs.DeCap .ButtonName }}HomeButton = &{{ .ButtonName }}HomeButton{}
)

// {{ .ButtonName }}HomeButton is a home button that switches to a tab bar panel when tapped.
type {{ .ButtonName }}HomeButton struct{}

func init() {
	builder.RegisterButton({{ .ButtonIndex }}, {{ call .Funcs.DeCap .ButtonName }}HomeButton)
}

func (b {{ .ButtonName }}HomeButton) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .ButtonName }}HomeButton.Init: %w", err)
		}
	}()

	var tabbarBuilder builder.TabBarContainerBuilder
	if tabbarBuilder = builder.TabBar({{ call .Funcs.DeCap .ButtonName }}ButtonName); tabbarBuilder == nil {
		err = fmt.Errorf("no tab bar builder for the button %q", {{ call .Funcs.DeCap .ButtonName }}ButtonName)
		return
	}
	err = tabbarBuilder.Init(ctx, ctxCancel, app, w)
	return
}

func (b {{ .ButtonName }}HomeButton) Button(backContent *fyne.Container) (button fyne.CanvasObject) {

	tabbarBuilder := builder.TabBar({{ call .Funcs.DeCap .ButtonName }}ButtonName)

	tappedContent := backpanel.Content(
		{{ call .Funcs.DeCap .ButtonName }}ButtonLabel,
		func() { window.SetContent(backContent) },
		tabbarBuilder.TabBar(),
	)

	button = safebutton.New(
		{{ call .Funcs.DeCap .ButtonName }}ButtonLabel,
		func() {
			window.SetContent(tappedContent)
		},
	)
	return
}

`
)
