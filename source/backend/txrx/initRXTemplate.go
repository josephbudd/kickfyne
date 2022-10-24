package txrx

type initRXTemplateData struct {
	ImportPrefix string
}

var initRXTemplate = `package txrx

import (
	"context"
	"fmt"

	"{{ .ImportPrefix }}/shared/message"
	"{{ .ImportPrefix }}/shared/store"
)

const initF = "receiveInit: %s"

func init() {
	addReceiver(message.InitID, receiveInit)
}

func receiveInit(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{}) {

	initMsg := msg.(*message.Init)
	var err, fatal error
	defer func() {
		switch {
		case err != nil:
			initMsg.Error = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, err.Error())
			Send(initMsg)
		case fatal != nil:
			initMsg.Fatal = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, fatal.Error())
			Send(initMsg)
		default:
			// No errors so don't send back the Init message.
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
