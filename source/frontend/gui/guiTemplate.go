package gui

const (
	guiFileName = "gui.go"
)

var guiTemplate = `package gui

import (
	"fyne.io/fyne/v2"
)

var window fyne.Window

// Init sets the app window var.
func Init(w fyne.Window) {
	window = w
}

// CanvasObjectProvider provides a screen's canvas object to ScreenCanvasWatchers.
// When a screen's canvas object changes, the new canvas object is passed in calls to the ScreenCanvasWatchers func Watch()
// A panel's func Show() changes it's screen's canvas object to that panel's canvas object.
type CanvasObjectProvider interface {
	BindToWindow()
	UpdateCanvasObject(canvasObject fyne.CanvasObject)
	CanvasObject() (canvasObject fyne.CanvasObject)
	BindWatcher(watcher ScreenCanvasWatcher)
	UnBindWatcher(watcher ScreenCanvasWatcher)
}

// ScreenCanvasWatcher sets a target's canvas object with an updated screen canvas object.
// For example:
// The type windowScreenCanvasObjectWatcher:
// * Updates the app's window with it's func Watch.
// * Use it for button's.
// * Use if for main menu navigation.
// The type TabItemScreenCanvasObjectWatcher:
// * Updates the TabItem's content with it's func Watch.
// * Use it for TabItems.
type ScreenCanvasWatcher interface {
	Watch(CanvasObjectProvider)
	UnBind()
}

`
