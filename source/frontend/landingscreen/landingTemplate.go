package landingscreen

type landingTemplateData struct {
	ImportPrefix      string
	LandingScreenName string
}

const (
	landingFileName = "landing.go"

	landingTemplate = `package landingscreen


import (
	"{{ .ImportPrefix }}/frontend/gui"
	"{{ .ImportPrefix }}/frontend/gui/screens/{{ .LandingScreenName }}"
)

var landingScreen gui.CanvasObjectProvider

// Init binds the landing screen to the app window.
func Init() {
	landingScreen = {{ .LandingScreenName }}.CanvasObjectProvider()
	// Show the landing screen right away even if it's not full of content.
	landingScreen.BindToWindow()
}

// Land binds the landing to the window.
func Land() {
	landingScreen.BindToWindow()
}

`
)
