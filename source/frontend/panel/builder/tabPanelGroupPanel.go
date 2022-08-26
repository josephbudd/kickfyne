package builder

const (
	tabPanelGroupPanelFileName = "tabPanelGroupPanel.go"

	tabPanelGroupPanelTemplate = `package builder

	// A panel builder.
	// home button name, tab name, builders.
	var tabPanelBuilderMap map[string]map[string][]PanelBuilder = make(map[string]map[string][]PanelBuilder)
	
	// RegisterTabPanel regiters a panel builder with a button name, tab name, panel group name.
	// Each panel must call this in func init().
	func RegisterTabPanel(homeButtonName, tabName string, panelBuilder PanelBuilder) {
		var found bool
		var tabNamesPanelBuilders map[string][]PanelBuilder
		if tabNamesPanelBuilders, found = tabPanelBuilderMap[homeButtonName]; !found {
			tabNamesPanelBuilders = make(map[string][]PanelBuilder)
			tabPanelBuilderMap[homeButtonName] = tabNamesPanelBuilders
		}
		var panelBuilders []PanelBuilder
		if panelBuilders, found = tabNamesPanelBuilders[tabName]; !found {
			panelBuilders = make([]PanelBuilder, 0, 5)
		}
		tabNamesPanelBuilders[tabName] = append(panelBuilders, panelBuilder)
	}
	
	// TabPanels returns group's panels.
	func TabPanels(homeButtonName, tabName string) (panelBuilders []PanelBuilder) {
		var found bool
		var tabNamesPanelBuilders map[string][]PanelBuilder
		if tabNamesPanelBuilders, found = tabPanelBuilderMap[homeButtonName]; !found {
			return
		}
		if panelBuilders, found = tabNamesPanelBuilders[tabName]; !found {
			panelBuilders = nil
		}
		return
	}

`
)
