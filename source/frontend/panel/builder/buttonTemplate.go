package builder

const (
	buttonFileName = "button.go"

	buttonTemplate = `package builder

	import (
		"sort"
	)
	
	// A home button and it's onclick view builder.
	
	var buttonMap map[int]ButtonBuilder = make(map[int]ButtonBuilder)
	
	// RegisterButton registers a home button builder with the home button pad.
	func RegisterButton(buttonIndex int, buttonBuilder ButtonBuilder) {
		var found bool
		if _, found = buttonMap[buttonIndex]; !found {
			buttonMap[buttonIndex] = buttonBuilder
			return
		}
		// Insert this before the existing item.
		// Get the keys and sort them.
		indexes := make([]int, 0, len(buttonMap))
		for index := range buttonMap {
			indexes = append(indexes, index)
		}
		sort.Ints(indexes)
		// Starting with the highest indexes, inc each one >= buttonIndex.
		last := len(indexes) - 1
		for i := last; i >= 0; i-- {
			oldIndex := indexes[i]
			// Inc this index to make room for buttonIndex.
			buttonMap[oldIndex+1] = buttonMap[oldIndex]
			if oldIndex != buttonIndex {
				// Indexes might not be in increments of 1 so remove the unwanted old item.
				delete(buttonMap, oldIndex)
			} else {
				// Add the builder.
				buttonMap[buttonIndex] = buttonBuilder
				return
			}
		}
	}
	
	// Buttons returns a sorted slice of home button builders.
	func Buttons() (buttonBuilders []ButtonBuilder) {
		buttonBuilders = make([]ButtonBuilder, 0, len(buttonMap))
		indexes := make([]int, 0, len(buttonMap))
		for index := range buttonMap {
			indexes = append(indexes, index)
		}
		sort.Ints(indexes)
		for _, index := range indexes {
			buttonViewBuilder := buttonMap[index]
			buttonBuilders = append(buttonBuilders, buttonViewBuilder)
		}
		return
	}

`
)
