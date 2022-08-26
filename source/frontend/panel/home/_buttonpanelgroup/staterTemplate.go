package buttonpanelgroup

var staterFileName = "stateHandler.go"

type staterTemplateData = struct {
	ImportPrefix string
	PackageName  string
}

var staterTemplate = `package {{ .PackageName }}

import (
	"{{ .ImportPrefix }}/shared/state"
)

var appState *state.FrontendState
var stater *stateHandler

type stateHandler struct{}

func (s *stateHandler) init() {
	appState = state.NewFrontendState()
	appState.AddListener(s)
}

// StateRX receives the message from the state.
func (s *stateHandler) StateRX(msg state.Message) {
	/* KICKFYNE TODO:
	Handle state changes indicated by msg in your panels.
	*/
}
`
