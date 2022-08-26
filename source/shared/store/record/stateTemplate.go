package record

const (
	stateFileName = "state.go"

	stateTemplate = `package record

type State struct {
	ID              uint64

	/* KICKFYNE TODO:
	Complete this State struct definition.
	*/
}

func NewState() (state State) {
	state = State{
		ID: 1,
	}
	return
}

// IsZero returns if the record hasn't been given an id.
// A record from the store has an ID > 0.
// Use to check the record returned by {{ .RecordName }}Store.Get(id)
func (state State) IsZero() (isZero bool) {
	isZero = state.ID == 0
	return
}

`
)
