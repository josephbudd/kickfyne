package selection

const (
	radioGroupFileName = "radioGroup.go"
)

var radioGroupTemplate = `package selection

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RadioGroup is a scrolling select list widget that is bound to it's data.
// In this case the data is []string.
type RadioGroup struct {
	radiogroup     *widget.RadioGroup
	options        []string
	content        fyne.CanvasObject
	onSelectedFunc func(index int)
}

// NewRadioGroup constructs a new RadioGroup with no data.
// Reboot([]string) resets the radio button options.
func NewRadioGroup(
	onSelected func(selectedIndex int),
) (rg *RadioGroup) {
	radiogroup := widget.NewRadioGroup(
		[]string{},
		nil,
	)

	rg = &RadioGroup{
		radiogroup:     radiogroup,
		options:        []string{},
		content:        container.NewMax(radiogroup),
		onSelectedFunc: onSelected,
	}
	radiogroup.OnChanged = rg.onSelected
	return
}

func (rg *RadioGroup) Reboot(options []string) {
	rg.options = options
	rg.radiogroup.Options = options
	rg.radiogroup.Refresh()
}

func (rg *RadioGroup) onSelected(selected string) {
	for i, s := range rg.options {
		if s == selected {
			rg.onSelectedFunc(i)
			return
		}
	}
}

// Widget returns the actual radio group widget.
func (rg *RadioGroup) Widget() (w *widget.RadioGroup) {
	w = rg.radiogroup
	return
}

`
