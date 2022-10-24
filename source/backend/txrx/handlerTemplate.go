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
	"{{ .ImportPrefix }}/shared/store"
)

const {{ $dCMessageName }}F = "receive{{ .MessageName }}: %s"

func init() {
	addReceiver(message.{{ .MessageName }}ID, receive{{ .MessageName }})
}

func receive{{ .MessageName }}(ctx context.Context, ctxCancel context.CancelFunc, stores *store.Stores, msg interface{}) {

	{{ $dCMessageName }}Msg := msg.(*message.{{ .MessageName }})
	var err, fatal error
	defer func() {
		switch {
		case fatal != nil:
			{{ $dCMessageName }}Msg.Fatal = true
			{{ $dCMessageName }}Msg.ErrorMessage = fmt.Sprintf({{ $dCMessageName }}F, fatal.Error())
			Send({{ $dCMessageName }}Msg)
		case err != nil:
			{{ $dCMessageName }}Msg.Error = true
			{{ $dCMessageName }}Msg.ErrorMessage = fmt.Sprintf({{ $dCMessageName }}F, err.Error())
			Send({{ $dCMessageName }}Msg)
		default:
			// No errors so return the {{ $dCMessageName }}Msg.
			Send({{ $dCMessageName }}Msg)
		}
	}()

	/* KICKFYNE TODO:
	Do something with this message.
	Use fatal for unrecoverable errors.
	Use err for user error messages.
	*/
}

`
