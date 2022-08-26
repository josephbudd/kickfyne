package txrx

const (
	tXRXFileName = "txrx.go"
)

type tXRXTemplateData struct {
	ImportPrefix string
}

var tXRXTemplate = `package txrx

import (
	"context"
	"fmt"
	"log"

	"{{ .ImportPrefix }}/shared/message"
	"{{ .ImportPrefix }}/shared/store"
)

type Listener func(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{})

var (
	messageListeners = make(map[uint64][]Listener, 20)
)

// addListener adds the number of listeners.
func addListener(msgID uint64, listener Listener) (err error) {
	if !message.IsValidID(msgID) {
		err = fmt.Errorf("message.AddListener: message id not found")
		return
	}
	var listeners []Listener
	var found bool
	if listeners, found = messageListeners[msgID]; !found {
		listeners = make([]Listener, 0, 20)
	}
	listeners = append(listeners, listener)
	messageListeners[msgID] = listeners
	return
}

// Listen listen form message from the main process and distributes them.
func Listen(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores) {
	go func(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Backend Listener DONE")
				return
			case msg := <-message.FrontEndToBackEnd:
				id := msg.ID()
				name := msg.Name()
				var listeners []Listener
				var found bool
				if listeners, found = messageListeners[id]; !found {
					log.Printf("backend listeners not found for *message.%s", name)
					continue
				}
				realMSG := msg.MSG()
				for _, l := range listeners {
					go l(ctx, ctxCancel, stores, realMSG)
				}
			}
		}
	}(ctx, ctxCancel, stores)
}

`
