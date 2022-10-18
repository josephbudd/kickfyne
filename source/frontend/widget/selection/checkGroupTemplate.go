package selection

const (
	checkGroupFileName = "checkGroup.go"
)

var checkGroupTemplate = `package selection

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CheckGroup is a scrolling select list widget that is bound to it's data.
// In this case the data is []string.
type CheckGroup struct {
	checks         []*widget.Check
	options        []string
	content        fyne.CanvasObject
	onSelectedFunc func(index int, checked bool)
}

// NewCheckGroup constructs a new CheckGroup with no data.
func NewCheckGroup(
	onSelected func(selectedIndex int, checked bool),
) (cg *CheckGroup) {
	cg = &CheckGroup{onSelectedFunc: onSelected}
	return
}

// Reboot([]string) resets the check box options.
func (cg *CheckGroup) Reboot(options []string) {
	cg.options = options
	cg.checks = make([]*widget.Check, len(options))
	canvasObjects := make([]fyne.CanvasObject, len(options))
	for i, s := range options {
		check := widget.NewCheck(
			s,
			func(checked bool) {
				cg.onSelectedFunc(i, checked)
			},
		)
		cg.checks[i] = check
		canvasObjects[i] = check
	}
	// Layout
	vbox := container.New(layout.NewVBoxLayout(), canvasObjects...)
	cg.content = container.NewMax(vbox)
	cg.content.Refresh()
}

// Widget returns the actual check group widget.
func (cg *CheckGroup) Widget() (w fyne.CanvasObject) {
	w = cg.content
	return
}

`
