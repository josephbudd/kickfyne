package builder

const (
	buttonPanelGroupFileName = "buttonPanelGroup.go"

	buttonPanelGroupTemplate = `package builder

	// Panel group builder for buttons.
	
	var buttonPanelGroupPanelsMap map[string]PanelGroupBuilder = make(map[string]PanelGroupBuilder)
	
	// RegisterButtonPanelGroup registers a panel group builder with a button name.
	func RegisterButtonPanelGroup(buttonName string, panelGroupBuilder PanelGroupBuilder) {
		buttonPanelGroupPanelsMap[buttonName] = panelGroupBuilder
	}
	
	// ButtonPanelGroup returns a button's panel group builder.
	func ButtonPanelGroup(buttonName string) (panelBuilder PanelGroupBuilder) {
		panelBuilder = buttonPanelGroupPanelsMap[buttonName]
		return
	}

`
)
