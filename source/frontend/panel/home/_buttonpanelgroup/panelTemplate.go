package buttonpanelgroup

import (
	"github.com/josephbudd/kickfyne/source/panelloader"
	"github.com/josephbudd/kickfyne/source/utils"
)

type panelTemplateData struct {
	PackageName string
	Panel       panelloader.Panel
	Funcs       utils.Funcs
}

var panelTemplate = `package {{ .PackageName }}

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
 	"fyne.io/fyne/v2/widget"
)

// {{ call .Funcs.DeCap .Panel.Name }}Panel
{{ call .Funcs.Comment .Panel.Description }}
var {{ call .Funcs.DeCap .Panel.Name }}Panel *{{ call .Funcs.DeCap .Panel.Name }}Components

// {{ call .Funcs.DeCap .Panel.Name }}Components contains each component and the content.
type {{ call .Funcs.DeCap .Panel.Name }}Components struct {

	heading *widget.Label
	content fyne.CanvasObject

	/* KICKFYNE TODO:
	Add other important components like labels, form inputs.

	// Example: A form.
	form *contactform.AddForm

	*/
}

func build{{ .Panel.Name }}Panel() {

	// Build each panel component.
	{{ call .Funcs.DeCap .Panel.Name }}Panel = &{{ call .Funcs.DeCap .Panel.Name }}Components{}
	{{ call .Funcs.DeCap .Panel.Name }}Panel.heading = widget.NewLabelWithStyle("{{ .Panel.Heading }}", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	/* KICKFYNE TODO:

	// Build the components that you added to the {{ call .Funcs.DeCap .Panel.Name }}Components struct.
	// Example: A form.

	{{ call .Funcs.DeCap .Panel.Name }}Panel.form = contactform.NewAdd(
		func(r *record.ContactAdd, err error) {
			if err != nil {
				dialog.ShowInformation("Error", err.Error(), window)
				return
			}
			messenger.contactAddTX(r)
		},
	)

	*/

	// Set the panel's content.
	// Use a different container if needed.
	{{ call .Funcs.DeCap .Panel.Name }}Panel.content = container.NewVBox(
		{{ call .Funcs.DeCap .Panel.Name }}Panel.heading,

		/* KICKFYNE TODO:
		Add your components to the container.
		// Example: A form.

		{{ call .Funcs.DeCap .Panel.Name }}Panel.form,

		*/
	)
}

`
