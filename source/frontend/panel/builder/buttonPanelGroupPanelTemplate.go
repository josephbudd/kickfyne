package builder

const (
	buttonPanelGroupPanelFileName = "buttonPanelGroupPanel.go"

	buttonPanelGroupPanelTemplate = `package builder

	// A panel builder.
	// home button name, builders.
	var buttonPanelBuilderMap map[string][]PanelBuilder = make(map[string][]PanelBuilder)
	
	// RegisterButtonPanel registers a panel builder with a button name.
	// Each panel must call this in func init().
	func RegisterButtonPanel(homeButtonName string, panelBuilder PanelBuilder) {
		var found bool
		var panelBuilders []PanelBuilder
		if panelBuilders, found = buttonPanelBuilderMap[homeButtonName]; !found {
			panelBuilders = make([]PanelBuilder, 0, 5)
		}
		buttonPanelBuilderMap[homeButtonName] = append(panelBuilders, panelBuilder)
	}
	
	// ButtonPanels returns group's panels.
	func ButtonPanels(homeButtonName string) (panelBuilders []PanelBuilder) {
		var found bool
		if panelBuilders, found = buttonPanelBuilderMap[homeButtonName]; !found {
			panelBuilders = nil
		}
		return
	}

`
)
