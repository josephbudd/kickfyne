package backpanel

const (
	fileName = "backpanel.go"
)

var template = `package backpanel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Content constructs the back panel.
// A back panel has a:
// * A heading.
// * A back button.
// * A vertically scrolling group content.
func Content(heading string, back func(), groupContent *fyne.Container) (content *fyne.Container) {
	label := widget.NewLabelWithStyle(heading, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	button := widget.NewButtonWithIcon(
		"Back",
		theme.NavigateBackIcon(),
		back,
	)
	scrollingGroupContent := container.NewVScroll(groupContent)
	content = container.NewBorder(label, nil, button, nil, scrollingGroupContent)
	return
}

`
