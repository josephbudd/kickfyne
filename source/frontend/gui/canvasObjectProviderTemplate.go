package gui

const (
	canvasObjectProviderFileName = "canvasObjectProvider.go"

	canvasObjectProviderTemplate = `package gui

import (
	"fyne.io/fyne/v2"
)

// ScreenCanvasManager passes a screen's canvas object to each watcher's func Watch, whenever the screen's canvas object changes.
// It is an implementation of the CanvasObjectProvider interface.
// Every screen has it's own ScreenCanvasManager.
type ScreenCanvasManager struct {
	canvasObject  fyne.CanvasObject
	windowWatcher ScreenCanvasWatcher
	watchers      []ScreenCanvasWatcher
}

// NewScreenCanvasManager constructs a new ScreenCanvasManager.
// Binds the windowContent to watching the canvas object.
func NewScreenCanvasManager() (screenCanvasManager *ScreenCanvasManager) {
	screenCanvasManager = &ScreenCanvasManager{
		watchers: make([]ScreenCanvasWatcher, 0, 5),
	}
	screenCanvasManager.windowWatcher = NewWindowScreenCanvasObjectWatcher()
	return
}

// BindToWindow makes the ScreenCanvasManager window watcher the current app window watcher.
func (screenCanvasManager *ScreenCanvasManager) BindToWindow() {
	SetCurrentWindowScreenCanvasObjectWatcher(screenCanvasManager.windowWatcher, screenCanvasManager.canvasObject)
}

// UpdateCanvasObject updates the canvas object and informs the watchers.
func (screenCanvasManager *ScreenCanvasManager) UpdateCanvasObject(canvasObject fyne.CanvasObject) {
	screenCanvasManager.canvasObject = canvasObject
	screenCanvasManager.windowWatcher.Watch(screenCanvasManager)
	for _, watcher := range screenCanvasManager.watchers {
		watcher.Watch(screenCanvasManager)
	}
}

// CanvasObject returns the CanvasObject.
// It will be called by the watcher.
func (screenCanvasManager *ScreenCanvasManager) CanvasObject() (canvasObject fyne.CanvasObject) {
	canvasObject = screenCanvasManager.canvasObject
	return
}

// BindWatcher adds a watcher to the ScreenCanvasManager.
// The ScreenCanvasManager will update the watcher whenever the screens canvas object is updated.
func (screenCanvasManager *ScreenCanvasManager) BindWatcher(watcher ScreenCanvasWatcher) {
	for _, w := range screenCanvasManager.watchers {
		if w == watcher {
			// Don't allow duplicates.
			return
		}
	}
	screenCanvasManager.watchers = append(screenCanvasManager.watchers, watcher)
	watcher.Watch(screenCanvasManager)
}

// UnBindWatcher removes a watcher to the ScreenCanvasManager.
func (screenCanvasManager *ScreenCanvasManager) UnBindWatcher(watcher ScreenCanvasWatcher) {
	watchers := screenCanvasManager.watchers
	for i, w := range watchers {
		if w == watcher {
			screenCanvasManager.watchers = watchers[0:i]
			if i++; i < len(watchers) {
				screenCanvasManager.watchers = append(screenCanvasManager.watchers, watchers[i:]...)
			}
			return
		}
	}
}

`
)
