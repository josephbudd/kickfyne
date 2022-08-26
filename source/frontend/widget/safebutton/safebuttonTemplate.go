package safebutton

const (
	fileName = "safebutton.go"
)

var template = `package safebutton

import (
	"fyne.io/fyne/v2/widget"
)

// SafeButton is a safe button that won't double click.
type SafeButton struct {
	widget.Button
}

// button the won't double tap.
func New(label string, ontap func()) (button *SafeButton) { // return type
	button = &SafeButton{
		widget.Button{
			Text: label,
			OnTapped: func() {
				button.Disable()
				ontap()
				button.Enable()
			},
		},
	}
	button.ExtendBaseWidget(button)
	return
}

`
