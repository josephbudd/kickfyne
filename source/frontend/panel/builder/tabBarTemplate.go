package builder

const (
	tabBarFileName = "tabBar.go"

	tabBarTemplate = `package builder

	// A tabbar builder.
	
	var tabbarBuilderMap map[string]TabBarContainerBuilder = make(map[string]TabBarContainerBuilder)
	
	// RegisterTabBar registers a tabbar builder with it's parent's group name.
	func RegisterTabBar(
		buttonName string,
		tabbarBuilder TabBarContainerBuilder,
	) {
		// Add the builder.
		tabbarBuilderMap[buttonName] = tabbarBuilder
	}
	
	// TabBar returns a sorted slice of tab builders.
	func TabBar(homeButtonName string) (tabbarBuilder TabBarContainerBuilder) {
		var found bool
		if tabbarBuilder, found = tabbarBuilderMap[homeButtonName]; !found {
			tabbarBuilder = nil
		}
		return
	}

`
)
