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

type Receiver func(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{})

var (
	messageReceivers = make(map[uint64][]Receiver, 20)
)

// Send sends a message to the front-end.
func Send(msg message.MSGer) {
	message.BackEndToFrontEnd <- msg
}

// addReceiver adds the number of receivers.
func addReceiver(msgID uint64, receiver Receiver) (err error) {
	if !message.IsValidID(msgID) {
		err = fmt.Errorf("message.AddReceiver: message id not found")
		return
	}
	var receivers []Receiver
	var found bool
	if receivers, found = messageReceivers[msgID]; !found {
		receivers = make([]Receiver, 0, 5)
	}
	receivers = append(receivers, receiver)
	messageReceivers[msgID] = receivers
	return
}

// StartReceiver starts receiving messages from the front-end and dispatches them to the back-end.
func StartReceiver(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores) {
	go func(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Backend Receiver DONE")
				return
			case msg := <-message.FrontEndToBackEnd:
				id := msg.ID()
				name := msg.Name()
				var receivers []Receiver
				var found bool
				if receivers, found = messageReceivers[id]; !found {
					log.Printf("backend receivers not found for *message.%s", name)
					continue
				}
				realMSG := msg.AsInterface()
				for _, f := range receivers {
					go f(ctx, ctxCancel, stores, realMSG)
				}
			}
		}
	}(ctx, ctxCancel, stores)
}

`
