package txrx

type templateData struct {
	ImportPrefix string
}

const (
	fileName = "listen.go"
)

var template = `package txrx

import (
	"context"
	"fmt"
	"log"

	"{{ .ImportPrefix }}/shared/message"
)

var messageListeners = make(map[uint64][]Listener, 20)
var Dispatcher = make(chan interface{}, 1024)

type Listener interface {
	Listen(msg interface{})
	GroupName() string
}

// AddListener adds the number of listeners.
func AddListener(msgID uint64, listener Listener) (err error) {
	if !message.IsValidID(msgID) {
		err = fmt.Errorf("message.AddListener: message id not found")
		return
	}
	var listeners []Listener
	var found bool
	if listeners, found = messageListeners[msgID]; !found {
		listeners = make([]Listener, 0, 20)
	}
	// Don't allow for duplicates.
	for _, l := range listeners {
		if l == listener {
			err = fmt.Errorf("%s is already listening to %d", listener.GroupName(), msgID)
			return
		}
	}
	listeners = append(listeners, listener)
	messageListeners[msgID] = listeners
	return
}

// RemoveListener removes the number of listeners.
func RemoveListener(msgID uint64, listener Listener) (err error) {
	if !message.IsValidID(msgID) {
		err = fmt.Errorf("message.RemoveListener: message id not found")
		return
	}
	if _, found := messageListeners[msgID]; !found {
		err = fmt.Errorf("message.RemoveListener: message already has no listeners")
		return
	}
	listeners := messageListeners[msgID]
	for i, l := range listeners {
		if l == listener {
			messageListeners[msgID] = listeners[0:i]
			messageListeners[msgID] = append(messageListeners[msgID], listeners[i:]...)
			return
		}
	}
	return
}

// Listen listen form message from the main process and distributes them.
func Listen(ctx context.Context, ctxCancel context.CancelFunc) {
	go func(ctx context.Context, ctxCancel context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Frontend Listener DONE")
				return
			case msg := <-message.BackEndToFrontEnd:
				var isFatal bool
				var errorMessage string
				if isFatal, errorMessage = msg.FatalError(); isFatal {
					log.Printf("frontend: txrx.Listen: Fatal from back end: %s", errorMessage)
					ctxCancel()
					return
				}
				id := msg.ID()
				var listeners []Listener
				var found bool
				if listeners, found = messageListeners[id]; !found {
					// No listeners for this message.
					continue
				}
				// Dispatch the message.
				realMSG := msg.MSG()
				for _, l := range listeners {
					go l.Listen(realMSG)
				}
			}
		}
	}(ctx, ctxCancel)
}

`
