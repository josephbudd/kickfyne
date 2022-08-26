package builder

const (
	tabPanelGroupFileName = "tabPanelGroup.go"

	tabPanelGroupTemplate = `package builder

	// Panel group builder for tabs.
	
	var tabPanelGroupPanelsMap map[string]map[string]PanelGroupBuilder = make(map[string]map[string]PanelGroupBuilder)
	
	// RegisterTabPanelGroup registers a panel group builder with a tab name.
	func RegisterTabPanelGroup(homeButtonName, tabName string, panelGroupBuilder PanelGroupBuilder) {
		var tabNamePanelGroupBuilderMap map[string]PanelGroupBuilder
		var found bool
		if tabNamePanelGroupBuilderMap, found = tabPanelGroupPanelsMap[homeButtonName]; !found {
			tabNamePanelGroupBuilderMap = make(map[string]PanelGroupBuilder)
			tabPanelGroupPanelsMap[homeButtonName] = tabNamePanelGroupBuilderMap
		}
		tabNamePanelGroupBuilderMap[tabName] = panelGroupBuilder
	}
	
	// TabPanelGroup returns  a tab's panel group builder.
	func TabPanelGroup(homeButtonName, tabName string) (panelGroupBuilder PanelGroupBuilder) {
		var tabNamePanelGroupBuilderMap map[string]PanelGroupBuilder
		var found bool
		if tabNamePanelGroupBuilderMap, found = tabPanelGroupPanelsMap[homeButtonName]; !found {
			return
		}
		panelGroupBuilder = tabNamePanelGroupBuilderMap[tabName]
		return
	}

`
)
