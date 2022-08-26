package state

const (
	frontendFileName = "frontend.go"

	frontendTemplate = `package state

type FrontendState struct{}

// NewFrontendState constructs a new state for the back end.
// The frontend state only reads from state.
// It allows panel group stateHandlers to listen for state messages. (see messenger.go)
func NewFrontendState() (feState *FrontendState) {
	feState = &FrontendState{}
	return
}

/* KICKFYNE TODO:
Add getters for the front end.

func (bestate BackendState) GetSomething() (something string) {
	lockState()
	appState.state.Something = something
	unlockState()
	return
}

*/

`
)
