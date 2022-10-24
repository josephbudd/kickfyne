package screens

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type panelTemplateData struct {
	PanelName    string
	PackageName  string
	AddShowFunc  bool
	ImportPrefix string
	Funcs        utils.Funcs
}

const (
	panelTemplate = `package {{ .PackageName }}

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// {{ .PanelName }}Components is a panel.
// KICKFYNE TODO: Correct this panel's doc comment.
type {{ .PanelName }}Components struct {
	content fyne.CanvasObject
	screen  *screenComponents

	/* KICKFYNE TODO:
	Add other important components like labels, form inputs.

	Example: A label.
	label *widget.Label

	*/
}

// new{{ call .Funcs.Cap .PanelName }}Components initializes this panel.
// It creates each component that makes up the panel.
// It uses the components to create the panel content.
// Returns the panel and the error.
func new{{ call .Funcs.Cap .PanelName }}Components(screen *screenComponents) (panel *{{ .PanelName }}Components, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.new{{ call .Funcs.Cap .PanelName }}: %w", err)
		}
	}()

	panel = &{{ .PanelName }}Components{
		screen: screen,
	}

	/* KICKFYNE TODO:
	Build the components that you added to the landingPanelComponents struct.

	Example: A label.
	panel.label = widget.NewLabel("Hello world")

	*/

	// Create the panel's content.
	// Use a different container if needed.
	vbox := container.NewVBox(

		/* KICKFYNE TODO:
		Add your components to the container.

		Example: A label.
		panel.label,

		*/

	)

	/* KICKFYNE TODO:
	You may want a different scroller.

	Example:
	// Scroll both ways.
	scroller := container.NewScroll(vbox)

	*/
	scroller := container.NewVScroll(vbox)
	panel.content = container.NewMax(scroller)
	return
}
{{- if .AddShowFunc }}

// show shows this panel and hides the others.
func (panel *{{ .PanelName }}Components) show() {
	panel.screen.canvasObjectProvider.UpdateCanvasObject(panel.content)
}
{{- end }}

`
)
