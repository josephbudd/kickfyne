package gui

const (
	accordionItemWatcherFileName = "accordionItemWatcher.go"

	accordionItemWatcherTemplate = `package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// AccordionItemScreenCanvasObjectWatcher implements ScreenCanvasWatcher.
type AccordionItemScreenCanvasObjectWatcher struct {
	accordionItem        *widget.AccordionItem
	canvasObjectProvider CanvasObjectProvider
}

// NewAccordionItemScreenCanvasObjectWatcher constructs a new AccordionItemScreenCanvasObjectWatcher.
func NewAccordionItemScreenCanvasObjectWatcher(label string, canvasObjectProvider CanvasObjectProvider) (accordionItem *widget.AccordionItem, watcher *AccordionItemScreenCanvasObjectWatcher) {
	accordionItem = widget.NewAccordionItem(
		label,
		canvasObjectProvider.CanvasObject(),
	)
	watcher = &AccordionItemScreenCanvasObjectWatcher{
		accordionItem:        accordionItem,
		canvasObjectProvider: canvasObjectProvider,
	}
	canvasObjectProvider.BindWatcher(watcher)
	return
}

// Watch updates the accordionItem content if needed.
// Watch is the implementation of ScreenCanvasWatcher.
func (watcher *AccordionItemScreenCanvasObjectWatcher) Watch(canvasObjectProvider CanvasObjectProvider) {
	var canvasObject fyne.CanvasObject
	if canvasObject = canvasObjectProvider.CanvasObject(); canvasObject == nil {
		return
	}
	if watcher.accordionItem.Detail == canvasObject {
		return
	}
	watcher.accordionItem.Detail = canvasObject
}

// UnBind stop the watcher from watching.
// UnBind is the implementation of ScreenCanvasWatcher.
func (watcher *AccordionItemScreenCanvasObjectWatcher) UnBind() {
	watcher.canvasObjectProvider.UnBindWatcher(watcher)
}

`
)
