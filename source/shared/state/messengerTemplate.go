package state

const (
	messsengerFileName = "messsenger.go"

	messsengerTemplate = `package state

import (
	"sync"
)

type Listener interface {
	StateRX(msg Message)
}

var listeners = make([]Listener, 0, 25)
var listenerLock sync.Mutex

// AddListener adds a front end listen to the state messages sent out by the back end.
func (feState *FrontendState) AddListener(l Listener) {

	listenerLock.Lock()
	defer listenerLock.Unlock()

	for _, listener := range listeners {
		if l == listener {
			// Already added.
			return
		}
	}
	listeners = append(listeners, l)
}

// RemoveListener removes a front end listener.
func (feState *FrontendState) RemoveListener(l Listener) {

	listenerLock.Lock()
	defer listenerLock.Unlock()

	for i, listener := range listeners {
		if l == listener {
			listeners = listeners[0:i]
			listeners = append(listeners, listeners[i:]...)
			return
		}
	}
}

// Dispatch dispatches the message to the front end listeners that the state has been updated.
func (bestate *BackendState) Dispatch(msg *Message) {

	listenerLock.Lock()
	defer listenerLock.Unlock()

	stateMsg := *msg
	for _, listener := range listeners {
		go listener.StateRX(stateMsg)
	}
}


`
)
