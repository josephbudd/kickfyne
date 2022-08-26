package state

const (
	messsageFileName = "messsage.go"

	messsageTemplate = `package state

// Message informs the front process of state changes.
type Message struct {

	/* KICKFYNE TODO:
	Complete the message definition with bool flags.

	Example:
	NewSomething bool
	*/
}

type StateWith func(msg *Message)

// NewMessage constructs a new State message.
func NewMessage(with ...StateWith) (msg *Message) {
	msg = &Message{}
	for _, w := range with {
		w(msg)
	}
	return
}

/* KICKFYNE TODO:
Add message-with funcs.

Example:
// StateWithNewSomething signals that the current something is new.
// The backend will use with NewState.
func StateWithNewSomething() (f func(msg *Message)) {
	f = func(msg *Message) {
		msg.NewSomething = true
	}
	return
}
*/

`
)
