package home

const (
	buttonPadFileName = "buttonPad.go"

	buttonPadTemplate = `package home

	import (
		"context"
		"fmt"
	
		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/container"
		"fyne.io/fyne/v2/layout"
	
		"{{ .ImportPrefix }}/frontend/panel/builder"
	)
	
	var (
		window fyne.Window
	)
	
	// Init initializes the panel groups.
	func Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("home.Init: %w", err)
			}
		}()
	
		window = w
	
		buttonBuilders := builder.Buttons()
		for _, buttonBuilder := range buttonBuilders {
			buttonBuilder.Init(ctx, ctxCancel, app, w)
		}
		return
	}
	
	// Content returns the current home content.
	func Content() (content *fyne.Container, err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("home.Content: %w", err)
			}
		}()
	
		content = container.NewCenter()
		buttonBuilders := builder.Buttons()
		objects := make([]fyne.CanvasObject, 0, len(buttonBuilders)*2)
		for i, buttonBuilder := range buttonBuilders {
			button := buttonBuilder.Button(content)
			if i > 0 {
				objects = append(objects, layout.NewSpacer())
			}
			objects = append(objects, button)
		}
		content.Add(
			container.NewHBox(objects...),
		)
		return
	}

`
)

type homeGroupTemplateData struct {
	ImportPrefix  string
	EffectImports []string
}
