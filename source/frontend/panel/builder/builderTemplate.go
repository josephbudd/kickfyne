package builder

const (
	builderFileName = "builder.go"

	builderTemplate = `package builder

	import (
		"context"
	
		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/container"
	)
	
	// ButtonBuilder creates content as a canvas object.
	type ButtonBuilder interface {
		Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error)
		Button(backContent *fyne.Container) (content fyne.CanvasObject)
	}
	
	// PanelGroupBuilder creates content as a canvas object.
	type PanelGroupBuilder interface {
		Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error)
		PanelGroup() (content fyne.CanvasObject)
	}
	
	// PanelBuilder creates content as a canvas object.
	type PanelBuilder interface {
		Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error)
		Panel() (content fyne.CanvasObject)
	}
	
	// TabBarContainerBuilder creates a tab bar content as a *fyne.Container.
	type TabBarContainerBuilder interface {
		Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error)
		TabBar() (tabbar *fyne.Container)
	}
	
	// TabItemBuilder is a tab item builder.
	type TabItemBuilder interface {
		Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error)
		TabItem() (tabItem *container.TabItem)
		Name() (name string)
	}

`
)
