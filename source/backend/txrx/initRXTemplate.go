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

const initF = "initRX: %s"

func init() {
	addListener(message.InitID, initRX)
}

func initRX(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{}) {

	initMsg := msg.(*message.Init)
	var err, fatal error
	defer func() {
		switch {
		case err != nil:
			initMsg.Error = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, err.Error())
			message.BackEndToFrontEnd <- initMsg
		case fatal != nil:
			initMsg.Fatal = true
			initMsg.ErrorMessage = fmt.Sprintf(initF, fatal.Error())
			message.BackEndToFrontEnd <- initMsg
		default:
			// No errors so don't send back the Init message.
		}
	}()

	/* KICKFYNE TODO:
	The front end is now ready to receive messages for initialization.
	Send those messages.
	*/

}

`
