package gui

const (
	windowWatcherFileName = "windowWatcher.go"

	windowWatcherTemplate = `package gui

import (
	"fyne.io/fyne/v2"
)

// currentWindowScreenCanvasObjectWatcher is the current watcher for the app window content.
var currentWindowScreenCanvasObjectWatcher ScreenCanvasWatcher

// windowScreenCanvasObjectWatcher implements ScreenCanvasWatcher
type windowScreenCanvasObjectWatcher struct {
	boundToWindow bool
}

// SetCurrentWindowScreenCanvasObjectWatcher sets the currentWindowScreenCanvasObjectWatcher to keep the screen content updated.
func SetCurrentWindowScreenCanvasObjectWatcher(watcher ScreenCanvasWatcher, screenCanvasObject fyne.CanvasObject) {
	currentWindowScreenCanvasObjectWatcher = watcher
	if window.Content() == screenCanvasObject {
		return
	}
	window.SetContent(screenCanvasObject)
}

// NewWindowScreenCanvasObjectWatcher constructs a new WindowScreenCanvasObjectWatcher.
func NewWindowScreenCanvasObjectWatcher() (watcher *windowScreenCanvasObjectWatcher) {
	watcher = &windowScreenCanvasObjectWatcher{}
	return
}

// Watch updates the window content if needed.
// Watch is the implementation of ScreenCanvasWatcher.
func (watcher *windowScreenCanvasObjectWatcher) Watch(canvasObjectProvider CanvasObjectProvider) {
	if currentWindowScreenCanvasObjectWatcher == watcher {
		var canvasObject fyne.CanvasObject
		if canvasObject = canvasObjectProvider.CanvasObject(); canvasObject == nil {
			return
		}
		if !watcher.boundToWindow {
			return
		}
		if window.Content() == canvasObject {
			return
		}
		window.SetContent(canvasObject)
	}
}

// UnBind stop the watcher from watching.
// It does not need to be used because it does nothing.
// UnBind is the implementation of ScreenCanvasWatcher.
func (watcher *windowScreenCanvasObjectWatcher) UnBind() {
}

`
)
