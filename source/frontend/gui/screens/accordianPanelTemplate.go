package screens

import "github.com/josephbudd/kickfyne/source/utils"

type accordionPanelTemplateData struct {
	PackageName  string
	PanelName    string
	ImportPrefix string
	Funcs        utils.Funcs
}

const (
	accordionPanelTemplate = `package {{ .PackageName }}

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"{{ .ImportPrefix }}/frontend/gui"
)

// {{ .PanelName }}Components is an Accordion panel.
// It is this screen's default panel.
type {{ .PanelName }}Components struct {
	content        fyne.CanvasObject
	screen         *screenComponents
	screenWatchers map[*widget.AccordionItem]*gui.AccordionItemScreenCanvasObjectWatcher
}

// new{{ call .Funcs.Cap .PanelName }} constructs this panel.
// It creates the accordion that makes up the panel.
// Returns the panel and error.
func new{{ call .Funcs.Cap .PanelName }}(screen *screenComponents) (panel *{{ .PanelName }}Components, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("{{ .PackageName }}.New{{ call .Funcs.Cap .PanelName }}: %w", err)
		}
	}()

	panel = &accordionPanelComponents{
		screen:         screen,
		screenWatchers: make(map[*widget.AccordionItem]*gui.AccordionItemScreenCanvasObjectWatcher),
	}

	// Build the accordionItems.
	var accordionItems []*widget.AccordionItem
	if accordionItems, err = panel.accordionItems(screen.ctx, screen.ctxCancel, screen.app, screen.window); err != nil {
		return
	}
	// Build the accordion.
	accordion := widget.NewAccordion(
		accordionItems...,
	)
	// Build the panel content.
	panel.content = container.New(
		layout.NewMaxLayout(),
		accordion,
	)
	return
}

func (panel *{{ .PanelName }}Components) accordionItems(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (items []*widget.AccordionItem, err error) {

	defer func() {
		if len(items) == 0 {
			err = fmt.Errorf("the Accordion in {{ .PackageName }}.{{ .PanelName }} must have items")
		}
		if err != nil {
			err = fmt.Errorf("panel.accordionItems: %w", err)
		}
	}()

	items = make([]*widget.AccordionItem, 0, 5)

	/* KICKFYNE TODO:
	Create each accordionItem for the Accordion widget.

	FOR EACH ACCORDION ITEM USING CONTENT FROM ANOTHER PANEL IN THIS SCREEN, use the 1 step example code below.
	// 1. Add the accordionItem.
	items = append(
		items,
		widget.NewAccordionItem("Inside", panel.screen.panels.insidePanel.content),
	)

	FOR EACH ACCORDION ITEM USING CONTENT FROM ANOTHER SCREEN, use the 2 step example code below.
	var otherScreen gui.CanvasObjectProvider
	// 1. Construct the other screen package.
	if otherScreen, err = outside.New(panel.screen.ctx, panel.screen.ctxCancel, panel.screen.app, panel.screen.window); err != nil {
		return
	}
	// 2. Build and add the accordionItem with the other screen's canvas object provider.
	items = append(
		items,
		panel.addScreenWatcherItem("Other Screen", otherScreen),
	)

	*/

	return
}

// addScreenWatcherItem creates and adds an AccordionItem with a canvas object provided by another screen.
func (panel *{{ .PanelName }}Components) addScreenWatcherItem(label string, otherScreen gui.CanvasObjectProvider) (accordionItem *widget.AccordionItem) {
	var watcher *gui.AccordionItemScreenCanvasObjectWatcher
	accordionItem, watcher = gui.NewAccordionItemScreenCanvasObjectWatcher(label, otherScreen)
	panel.screenWatchers[accordionItem] = watcher
	return
}

// Show shows this panel and hides the others.
func (panel *{{ .PanelName }}Components) Show() {
	panel.screen.canvasObjectProvider.UpdateCanvasObject(panel.content)
}

`
)
