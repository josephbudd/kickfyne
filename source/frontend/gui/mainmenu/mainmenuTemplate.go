package mainmenu

import "github.com/josephbudd/kickfyne/source/utils"

type mainMenuTemplateData struct {
	ImportPrefix       string
	ScreenPackageNames []string
	HomePackageName    string
	Funcs              utils.Funcs
}

const (
	mainMenuFileName = "mainmenu.go"

	mainMenuTemplate = `{{ $DOT := . }}package mainmenu

import(
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	"{{ .ImportPrefix }}/frontend/gui"
{{- range $screenName := .ScreenPackageNames }}
 {{- if ne $screenName $DOT.HomePackageName }}
	"{{ $DOT.ImportPrefix }}/frontend/gui/screens/{{ $screenName }}"
 {{- end }}
{{- end }}
	"{{ .ImportPrefix }}/frontend/landingscreen"
	"{{ .ImportPrefix }}/shared/meta"
)

var (
	appName       string
	label         string
	window        fyne.Window
	ctxCancelFunc context.CancelFunc
	menu          *fyne.Menu
)

// Init builds the main menu and adds it to the app.
func Init(ctx context.Context, cCFunc context.CancelFunc, app fyne.App, w fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.Init: %w", err)
		}
	}()

	// Build the menu.
	var data meta.AppData
	if data, err = meta.Data(); err != nil {
		return
	}
	appName = data.Details.Name
	label = appName
	window = w
	ctxCancelFunc = cCFunc
	return
}

// Reset reconstructs the main menu with items.
// If param items is nil it will be replaced using menuItems().
func Reset(items []*fyne.MenuItem) {
	allItems := make([]*fyne.MenuItem, 1, 0+2)
	// Prefix with Home item.
	item := &fyne.MenuItem{Label: "{{ call .Funcs.Cap .HomePackageName }}"}
	item.Action = func() {
		landingscreen.Land()
		resetLabel(appName)
	}
	allItems[0] = item
	// Add the customItems.
	if items == nil {
		items = menuItems()
	}
	if len(items) > 0 {
		allItems = append(allItems, items...)
	}
	// Suffix with Quit item if not a mobile device.
	if !fyne.CurrentDevice().IsMobile() {
		item = &fyne.MenuItem{
			Label:  "Quit",
			Action: ctxCancelFunc,
		}
		allItems = append(allItems, item)
	}
	menu = fyne.NewMenu(label, allItems...)
	window.SetMainMenu(
		fyne.NewMainMenu(
			menu,
		),
	)
}

func menuItems() (items []*fyne.MenuItem) {
{{- if ne (len .ScreenPackageNames) 0 }}
	items = make([]*fyne.MenuItem, 0, {{ len .ScreenPackageNames }})
{{- else }}
	items = make([]*fyne.MenuItem, 0)
{{- end }}
	/* KICKFYNE TODO:
	Create each item for the main menu.

	Example:
	var item *fyne.MenuItem
	// SomeScreen.
	// 1. Build the item.
	item = newMenuItem("SomeScreen", "SomeScreen - Do Something.", somescreen.CanvasObjectProvider())
	// 2. Add the item.
	items = append(items, item)

	*/

	return
}

func newMenuItem(label, heading string, screen gui.CanvasObjectProvider) (item *fyne.MenuItem) {
	item = &fyne.MenuItem{Label: label}
	item.Action = func() {
		// Bind the screen's content to the window so it is displayed.
		screen.BindToWindow()
		// Reset the menu label to the menu item's label.
		resetLabel(heading)
	}
	return
}

// resetLabel resets the label above the main menu.
func resetLabel(newLabel string) {
	menu.Label = newLabel
	menu.Refresh()
}

`
)
