package mainmenu

type mainMenuTemplateData struct {
	ImportPrefix      string
	LandingScreenName string
}

const (
	mainMenuFileName = "mainmenu.go"

	mainMenuTemplate = `package mainmenu

import(
	"context"
	"fmt"

	"fyne.io/fyne/v2"

	"{{ .ImportPrefix }}/frontend/landingscreen"
	"{{ .ImportPrefix }}/shared/metadata"
)

// Init builds the main menu and adds it to the app.
func Init(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("mainmenu.Init: %w", err)
		}
	}()

	// Build all the menu items.
	var items []*fyne.MenuItem
	if items, err = menuItems(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}

	// Build the menu.
	var data fyne.AppMetadata
	if data, err = metadata.AppMetaData(); err != nil {
		return
	}

	menu := fyne.NewMenu(
		data.Name,
		items...,
	)
	mainmenu := fyne.NewMainMenu(menu)
	window.SetMainMenu(mainmenu)
	return
}

func menuItems(ctx context.Context, ctxCancelFunc context.CancelFunc, app fyne.App, window fyne.Window) (items []*fyne.MenuItem, err error) {

	var item *fyne.MenuItem
	items = make([]*fyne.MenuItem, 0, 5)
	defer func() {
		// Add quit if not a mobile device.
		if !fyne.CurrentDevice().IsMobile() {
			item := fyne.NewMenuItem(
				"Quit",
				ctxCancelFunc,
			)
			items = append(items, item)
		}
	}()

	// Landing.
	item = fyne.NewMenuItem("Home", landingscreen.Land)
	items = append(items, item)

	/* KICKFYNE TODO:
	Create each item for the main menu.

	Example:
	var screen gui.CanvasObjectProvider
	// MyScreen.
	// 1. Get the screen.
	if screen, err = myscreen.New(ctx, ctxCancelFunc, app, window); err != nil {
		return
	}
	// 2. Build the item.
	item = fyne.NewMenuItem("MyScreen", screen.BindToWindow)
	// 3. Add the item.
	items = append(items, item)

	*/

	return
}

`
)
