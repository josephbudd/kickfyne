package buttontabbar

import "github.com/josephbudd/kickfyne/source/utils"

type staterTemplateData = struct {
	ImportPrefix   string
	PackageName    string
	HomeButtonName string
	Funcs          utils.Funcs
}

const (
	staterFileName = "stateHandler.go"

	staterTemplate = `package {{ .PackageName }}

import (
	"{{ .ImportPrefix }}/shared/state"
)

var appState *state.FrontendState
var stater *stateHandler = &stateHandler{}

type stateHandler struct{}

func (s *stateHandler) listen() {
	appState = state.NewFrontendState()
	appState.AddListener(stater)
}

// StateRX receives the message from the state.
func (s *stateHandler) StateRX(msg state.Message) {
	/* KICKFYNE TODO:
	Changes in state may or maynot effect whith tab is selected in this tab bar.
	*/
}

`
)
