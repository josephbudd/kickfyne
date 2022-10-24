package gui

const (
	tabItemWatcherFileName = "tabItemWatcher.go"

	tabItemWatcherTemplate = `package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type TabSelector interface {
	Select(*container.TabItem)
}

// TabItemScreenCanvasObjectWatcher implements ScreenCanvasWatcher.
type TabItemScreenCanvasObjectWatcher struct {
	tabItem              *container.TabItem
	canvasObjectProvider CanvasObjectProvider
	tabbar               TabSelector
}

// NewTabItemScreenCanvasObjectWatcher constructs a new TabItemScreenCanvasObjectWatcher.
func NewTabItemScreenCanvasObjectWatcher(label string, canvasObjectProvider CanvasObjectProvider, tabbar TabSelector) (tabItem *container.TabItem, watcher *TabItemScreenCanvasObjectWatcher) {
	tabItem = container.NewTabItem(
		label,
		canvasObjectProvider.CanvasObject(),
	)
	watcher = &TabItemScreenCanvasObjectWatcher{
		tabItem:              tabItem,
		canvasObjectProvider: canvasObjectProvider,
		tabbar:               tabbar,
	}
	canvasObjectProvider.BindWatcher(watcher)
	return
}

// Watch updates the tabItem content if needed.
// Watch is the implementation of ScreenCanvasWatcher.
func (watcher *TabItemScreenCanvasObjectWatcher) Watch(canvasObjectProvider CanvasObjectProvider) {
	var canvasObject fyne.CanvasObject
	if canvasObject = canvasObjectProvider.CanvasObject(); canvasObject == nil {
		return
	}
	watcher.tabItem.Content = canvasObject
	watcher.tabbar.Select(watcher.tabItem)
}

// UnBind stop the watcher from watching.
// UnBind is the implementation of ScreenCanvasWatcher.
func (watcher *TabItemScreenCanvasObjectWatcher) UnBind() {
	watcher.canvasObjectProvider.UnBindWatcher(watcher)
}

`
)
