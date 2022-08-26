package tabpanelgroup

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type panelTemplateData struct {
	PanelName        string
	PanelDescription string
	PanelHeading     string
	PackageName      string
	ImportPrefix     string
	Funcs            utils.Funcs
}

const panelTemplate = `package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"{{ .ImportPrefix }}/frontend/panel/builder"
)

// {{ call .Funcs.DeCap .PanelName }}Panel
{{ call .Funcs.Comment .PanelDescription }}
var {{ call .Funcs.DeCap .PanelName }}Panel *{{ call .Funcs.DeCap .PanelName }}PanelComponents = &{{ call .Funcs.DeCap .PanelName }}PanelComponents{}

func init() {
	builder.RegisterTabPanel(homeButtonName, tabName, {{ call .Funcs.DeCap .PanelName }}Panel)
}

type {{ call .Funcs.DeCap .PanelName }}PanelComponents struct {
	heading *widget.Label
	content fyne.CanvasObject

	/* KICKFYNE TODO:

	Add other important components like labels, form inputs.

	// Example: A form.
	form *contactform.AddForm

	*/
}

// Init .
// Init initializes this panel.
// It creates each component that makes up the panel.
// It uses the components to create the panel content.
// Returns the error.
func (p *{{ call .Funcs.DeCap .PanelName }}PanelComponents) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ call .Funcs.DeCap .PanelName }}PanelComponents.Init: %w", err)
		}
	}()

	// Build each panel component.
	{{ call .Funcs.DeCap .PanelName }}Panel.heading = widget.NewLabelWithStyle("{{ .PanelHeading }}", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	/* KICKFYNE TODO:

	// Build the components that you added to the {{ call .Funcs.DeCap .PanelName }}PanelComponents struct.
	// Example: A form.

	{{ call .Funcs.DeCap .PanelName }}Panel.form = contactform.NewAdd(
		func(r *record.ContactAdd, err error) {
			if err != nil {
				dialog.ShowInformation("Error", err.Error(), window)
				return
			}
			messenger.contactAddTX(r)
		},
	)

	*/

	// Create the panel's content.
	// Use a different container if needed.
	{{ call .Funcs.DeCap .PanelName }}Panel.content = container.NewVBox(
		{{ call .Funcs.DeCap .PanelName }}Panel.heading,

		/* KICKFYNE TODO:
		Add your components to the container.
		// Example: A form.

		{{ call .Funcs.DeCap .PanelName }}Panel.form,

		*/
	)
	return
}

// Panel returns this panel's content.
func (p *{{ call .Funcs.DeCap .PanelName }}PanelComponents) Panel() (content fyne.CanvasObject) {
	content = p.content
	return
}

func (p *{{ call .Funcs.DeCap .PanelName }}PanelComponents) Show() {
	// Get the panel builders.
	panelBuilders := builder.TabPanels(homeButtonName, tabName)
	for _, panelBuilder := range panelBuilders {
		if panelBuilder == p {
			p.content.Show()
		} else {
			content := panelBuilder.Panel()
			content.Hide()
		}
	}
}

`
