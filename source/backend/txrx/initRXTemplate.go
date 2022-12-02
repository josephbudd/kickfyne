package txrx

type initRXTemplateData struct {
	ImportPrefix string
}

var initRXTemplate = `package txrx

import (
	"context"
	"fmt"

	"{{ .ImportPrefix }}/shared/message"
	"{{ .ImportPrefix }}/backend/store"
)

func init() {
	addReceiver(message.InitID, receiveInit)
}

func receiveInit(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{}) {

	initMsg := msg.(*message.Init)
	var fatal error
	defer func() {
		if fatal != nil {
			initMsg.Fatal = true
			initMsg.ErrorMessage = fmt.Sprintf("receiveInit: %s", fatal.Error())
			Send(initMsg)
		}
	}()

	/* KICKFYNE TODO:
	
	1. The GUI just got displayed.
	   If parts of the back-end need initialized then initialize them now.
	2. The front-end screens are ready to receive messages for initialization.
	   Send those messages now.
	*/

}

`
