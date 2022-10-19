package screens

import "github.com/josephbudd/kickfyne/source/utils"

type doctabsPanelData struct {
	PackageName  string
	PanelName    string
	ImportPrefix string
	Funcs        utils.Funcs
}

const (
	doctabsPanelTemplate = `package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"{{ .ImportPrefix }}/frontend/gui"
)

// {{ .PanelName }}Components is an DocTabs panel.
// It is this screen's default panel.
type {{ .PanelName }}Components struct {
	content        fyne.CanvasObject
	screen         *screenComponents
	screenWatchers map[*container.TabItem]*gui.TabItemScreenCanvasObjectWatcher
}

// new{{ call .Funcs.Cap .PanelName }}Components constructs this panel.
// It creates the DocTabs container that makes up the panel.
// Returns the error.
func new{{ call .Funcs.Cap .PanelName }}Components(screen *screenComponents) (panel *{{ .PanelName }}Components, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.new{{ call .Funcs.Cap .PanelName }}: %w", err)
		}
	}()

	// Build the tabItems.
	var tabItems []*container.TabItem
	if tabItems, err = panel.tabItems(ctx, ctxCancel, app, w); err != nil {
		return
	}
	// Build the DocTabs container.
	apptabs := container.NewDocTabs(
		tabItems...,
	)
	// Build the panel content.
	panel.content = container.New(
		layout.NewMaxLayout(),
		apptabs,
	)
	return
}

func (panel *{{ .PanelName }}Components) tabItems(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (items []*container.TabItem, err error) {

	defer func() {
		if len(items) == 0 {
			err = fmt.Errorf("The DocTabs in {{ .PackageName }}.{{ .PanelName }} must have items.")
		}
		if err != nil {
			err = fmt.Errorf("panel.tabItems: %w", err)
		}
	}()

	items = make([]*container.TabItem, 0, 5)

	/* KICKFYNE TODO:
	Create each tabItem for the DocTabs container.

	FOR EACH TAB ITEM USING CONTENT FROM ANOTHER PANEL IN THIS SCREEN, use the 1 step example code below.
	// 1. Add the tabItem.
	items = append(
		items,
		container.NewTabItem("Inside", panel.screen.panels.insidePanel.content),
	)

	FOR EACH TAB ITEM USING CONTENT FROM ANOTHER SCREEN, use the 2 step example code below.
	var otherScreen gui.CanvasObjectProvider
	// 1. Construct the other screen package that provides the canvas object.
	if otherScreen, err = outside.New(panel.screen.ctx, panel.screen.ctxCancel, panel.screen.app, panel.screen.window); err != nil {
		return
	}
	// 2. Build and add the TabItem with the other screen's canvas object provider.
	items = append(
		items,
		panel.addScreenWatcherItem("Other Screen", otherScreen),
	)

	*/

	return
}

// addScreenWatcherItem creates and adds a TabItem with a canvas object provided by another screen.
func (panel *{{ .PanelName }}Components) addScreenWatcherItem(label string, otherScreen gui.CanvasObjectProvider) (accordionItem *widget.AccordionItem) {
	var watcher *gui.TabItemScreenCanvasObjectWatcher
	tabItem, watcher = gui.NewTabItemScreenCanvasObjectWatcher("Other Screen", otherScreen.CanvasObjectProvider())
	panel.screenWatchers[tabItem] = watcher
	return
}

// Show shows this panel and hides the others.
func (panel *{{ .PanelName }}Components) Show() {
	panel.screen.canvasObjectProvider.UpdateCanvasObject(panel.content)
}

`
)