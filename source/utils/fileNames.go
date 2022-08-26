package utils

func TabFileName(tabName string) (fileName string) {
	fileName = DeCap(tabName) + tabFileSuffix
	return
}

func PanelFileName(panelName string) (fileName string) {
	fileName = DeCap(panelName) + PanelFileSuffix
	return
}

func ButtonFileName(buttonName string) (fileName string) {
	fileName = DeCap(buttonName) + buttonFileSuffix
	return
}
