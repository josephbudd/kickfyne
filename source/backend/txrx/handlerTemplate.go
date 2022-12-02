package txrx

import "github.com/josephbudd/kickfyne/source/utils"

type handlerTemplateData struct {
	ImportPrefix string
	MessageName  string
	Funcs        utils.Funcs
}

var handlerTemplate = `{{ $dCMessageName := call .Funcs.DeCap .MessageName }}package txrx

import (
	"context"
	"fmt"

	"{{ .ImportPrefix }}/shared/message"
	"{{ .ImportPrefix }}/backend/store"
)

func init() {
	addReceiver(message.{{ .MessageName }}ID, receive{{ .MessageName }})
}

func receive{{ .MessageName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{}) {

	{{ $dCMessageName }}Msg := msg.(*message.{{ .MessageName }})
	var fatal error
	var userMessage string
	defer func() {
		switch {
		case fatal != nil:
			{{ $dCMessageName }}Msg.Fatal = true
			{{ $dCMessageName }}Msg.ErrorMessage = fmt.Sprintf("receive{{ .MessageName }}: %s", fatal.Error())
			Send({{ $dCMessageName }}Msg)
		case len(userMessage) > 0:
			{{ $dCMessageName }}Msg.Error = true
			{{ $dCMessageName }}Msg.ErrorMessage = userMessage
			Send({{ $dCMessageName }}Msg)
		default:
			// No errors so return the {{ $dCMessageName }}Msg.
			Send({{ $dCMessageName }}Msg)
		}
	}()

	/* KICKFYNE TODO:
	Do something with this message.
	Use fatal for unrecoverable errors.
	Use userMessage for user error messages.
	*/
}

`
