package builder

const (
	tabFileName = "tab.go"

	tabTemplate = `package builder

	import (
		"sort"
	)
	
	// A tab builder.
	
	var tabItemBuilderMap map[string]map[int]TabItemBuilder = make(map[string]map[int]TabItemBuilder)
	
	// RegisterTabItem registers a tab builder with it's tabbar's group name.
	// Each tab's panel group must call this in func init().
	func RegisterTabItem(
		homeButtonName string,
		tabIndex int,
		tabItemBuilder TabItemBuilder,
	) {
		var indexBuilderMap map[int]TabItemBuilder
		var found bool
		if indexBuilderMap, found = tabItemBuilderMap[homeButtonName]; !found {
			indexBuilderMap = make(map[int]TabItemBuilder)
			tabItemBuilderMap[homeButtonName] = indexBuilderMap
		}
		if _, found = indexBuilderMap[tabIndex]; !found {
			indexBuilderMap[tabIndex] = tabItemBuilder
			return
		}
		// Insert before an existing builder in the map.
		// Get the keys and sort them.
		indexes := make([]int, 0, len(indexBuilderMap))
		for index := range indexBuilderMap {
			indexes = append(indexes, index)
		}
		sort.Ints(indexes)
		// Starting with the highest indexes, inc each one >= tabIndex.
		last := len(indexes) - 1
		for i := last; i >= 0; i-- {
			oldIndex := indexes[i]
			// Inc this index to make room for tabIndex.
			indexBuilderMap[oldIndex+1] = indexBuilderMap[oldIndex]
			if oldIndex != tabIndex {
				// Indexes might not be in increments of 1 so remove the unwanted old item.
				delete(indexBuilderMap, oldIndex)
			} else {
				// Add the builder.
				indexBuilderMap[tabIndex] = tabItemBuilder
				return
			}
		}
	}
	
	// TabItems a tabbar's tab item builders in sorted order.
	func TabItems(homeButtonName string) (tabItemBuilders []TabItemBuilder) {
		var found bool
		var indexBuilderMap map[int]TabItemBuilder
		if indexBuilderMap, found = tabItemBuilderMap[homeButtonName]; !found {
			return
		}
		indexes := make([]int, 0, len(indexBuilderMap))
		for index := range indexBuilderMap {
			indexes = append(indexes, index)
		}
		sort.Ints(indexes)
		tabItemBuilders = make([]TabItemBuilder, 0, len(indexes))
		for _, index := range indexes {
			tabItemBuilders = append(tabItemBuilders, indexBuilderMap[index])
		}
		return
	}

`
)
