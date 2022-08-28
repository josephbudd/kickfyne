package buttontabbar

import (
	"github.com/josephbudd/kickfyne/source/utils"
)

type tabBarTemplateData struct {
	ImportPrefix   string
	PackageName    string
	HomeButtonName string
	Funcs          utils.Funcs
}

const (
	tabBarFileName = "tabBar.go"

	tabBarTemplate = `package {{ .PackageName }}

	import (
		"context"
		"fmt"
	
		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/container"
		"fyne.io/fyne/v2/layout"
		"{{ .ImportPrefix }}/frontend/panel/builder"
	)
	
	const (
		homeButtonName = "{{ .HomeButtonName }}"
	)
	
	var (
		{{ call .Funcs.DeCap .HomeButtonName }}TabBar *{{ .HomeButtonName }}TabBar = &{{ .HomeButtonName }}TabBar{}
		tabBar           *container.AppTabs
	)
	
	type {{ .HomeButtonName }}TabBar struct{}
	
	func init() {
		builder.RegisterTabBar(homeButtonName, {{ call .Funcs.DeCap .HomeButtonName }}TabBar)
	}
	
	// Implement the builder.TabBarContainerBuilder interface with funcs Init and TabBar.
	
	// Init does what is necessisary to initialize this view.
	// This is a tabbar group so init must Init each tab's view.
	// Returns the error.
	func (bldr *{{ .HomeButtonName }}TabBar) Init(ctx context.Context, ctxCancel context.CancelFunc, app fyne.App, w fyne.Window) (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("{{ .HomeButtonName }}TabBar.Init: %w", err)
			}
		}()
	
		builders := builder.TabItems(homeButtonName)
		for _, b := range builders {
			if b.Init(ctx, ctxCancel, app, w); err != nil {
				return
			}
		}
	
		// Start the messenger so it's communicating with the back end.
		err = messenger.listen()
		return
	}
	
	// TabBar creates and then returns this view's content as a container.
	func (bldr *{{ .HomeButtonName }}TabBar) TabBar() (tabbar *fyne.Container) {
		builders := builder.TabItems(homeButtonName)
		tabItems := make([]*container.TabItem, len(builders))
		for i, b := range builders {
			tabItem := b.TabItem()
			tabItems[i] = tabItem
		}
		tabBar = container.NewAppTabs(tabItems...)
		tabbar = container.New(
			layout.NewMaxLayout(),
			tabBar,
		)
		return
	}

`
)
