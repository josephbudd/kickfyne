package state

const (
	backendFileName = "backend.go"

	backendTemplate = `package state

type BackendState struct{}

// NewBackendState constructs a new state for the back end.
// The backend state writes to and reads from state.
// It also saves the state when it writes to the state.
// It also dispatches state messages to the front end when it writes to the state. (see messenger.go)
func NewBackendState() (beState BackendState) {
	beState = BackendState{}
	return
}


/* KICKFYNE TODO:

func (bestate BackendState) SetSomething(something string) (msg *Message, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("BackendState.SetSomething: %w", err)
		}
	}()

	lockState()
	defer unlockState()
	appState.state.Something = something
	if err = appState.stores.State.Update(appState.state); err != nil {
		unlockState()
		return
	}

	unlockState()
	msg = NewMessage(StateWithNewSomething())
	return
}

*/

`
)
